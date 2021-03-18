package proxy

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/koding/websocketproxy"
)

const (
	Middle_Layer_Proxy = iota + 1 //中间层
	User_Layer_Proxy              //外层，用户层

)
const (
	SecureBy_None = iota + 1 // 不需要验证用户
	SecureBy_User            //验证用户 采用token验证

)

const (
	Protocol_TCP = iota + 1
	Protocol_UDP
	Protocol_HTTP
	Protocol_HTTPS
	Protocol_SSH
	Protocol_WS
	Protocol_WSS
	Protocol_GRPC
)

const (
	TokenName  = "Secret-Token"
	BaseDomain = "z-gour.com" //配置基础域名
)

type ServerOptions struct {
	HTTPPort      int
	HTTPSPort     int
	SSHPort       int
	BaseDomain    string
	ServerLayer   int      // 代理运行层次，外层代理或中间层代理
	ForwardServer struct { //转发中间层代理服务配置
		Existed bool   //是否存在中间层代理
		IP      string //转发中间层代理IP
		Port    string //转发中间层代理端口
	}
	SSLCertFile string
	SSLKeyFile  string
}

//需要代理的网络目标请求，可以是HTTP 请求，HTTPS请求以及Websocket请求等
type TargetRequest struct {
	ID         string `json:"id"`
	UserName   string //用户名称
	RequestURL string //泛域名请求URL
	Protocol    int    // 代理协议
	Domain      string //外部域名
	SSLCertFile []byte //ssl 证书
	SSLKeyFile  []byte //ssl 证书
	TargetIP    string //目标IP
	TargetPort  string //目标端口
	//TargetProtocl string
	SecuryBy int // 安全配置

}

// todo  https://github.com/golang/go/issues/26479
//泛域名代理服务器
type PanDomainServer struct {
	targetRequest  *TargetRequest
	httpProxy      *httputil.ReverseProxy
	httpsProxy     *httputil.ReverseProxy
	webSocketProxy *websocketproxy.WebsocketProxy
}

var (
	httpServerListener  net.Listener
	httpsServerListener net.Listener
	sshServerListener   net.Listener
	errStream           chan error //错误信息

)

var (
	serverOpts ServerOptions
	//泛域名服务器，key为用户泛域名前缀，value为泛域名代理服务
	//example pro.z-gour.com;--->map['pro']=&pandomainserver
	domainServers = make(map[string]*PanDomainServer)
)

func init() {
}

//根据代理

func ServerInitWithOps(serverOpts ServerOptions) error {
	var (
		err          error
		webRemoteStr string
		webremote    *url.URL
		domainServer = &PanDomainServer{
			targetRequest: &TargetRequest{
				ID:         "",
				UserName:   "",
				RequestURL: "",

				Protocol:    Protocol_HTTP,
				Domain:      "",
				SSLCertFile: nil,
				SSLKeyFile:  nil,
				TargetIP:    "127.0.0.1",
				TargetPort:  "8000",
				SecuryBy:    0,
			},
			//httpProxy:      httputil.NewSingleHostReverseProxy(webremote),
			//httpsProxy:     httputil.NewSingleHostReverseProxy(webremote),
			webSocketProxy: nil,
		}
	)

	switch serverOpts.ServerLayer {

	case User_Layer_Proxy:
		if serverOpts.ForwardServer.Existed { //存在中间层代理
			switch domainServer.targetRequest.Protocol {
			case Protocol_HTTP: //进行转发，将请求转发到中间代理层
				webRemoteStr = "http://" + serverOpts.ForwardServer.IP + ":" + serverOpts.ForwardServer.Port + "/"
			}
		} else { //不存在中间层代理，则用户层代理直接对接底层目标服务器，即只有单层代理
			switch domainServer.targetRequest.Protocol { //对代理协议进行处理
			case Protocol_HTTP:
				webRemoteStr = "http://" + domainServer.targetRequest.TargetIP + ":" + domainServer.targetRequest.TargetPort
			case Protocol_HTTPS:
			//todo 处理https请求
			case Protocol_WS:
				//todo 处理websocket

			default:
				return nil
			}
		}

	case Middle_Layer_Proxy:
		if serverOpts.ForwardServer.Existed { //存在中间层，多层中间代理层情况 ，在这里就进行请求转发，将请求转发到下一层的中间代理服务中
			switch domainServer.targetRequest.Protocol {
			case Protocol_HTTP: //进行转发，将请求转发到中间代理层
				webRemoteStr = "http://" + serverOpts.ForwardServer.IP + ":" + serverOpts.ForwardServer.Port + "/"
			case Protocol_HTTPS:
			//todo 处理https请求
			case Protocol_WS:
				//todo 处理websocket

			}

		} else { //只有一层中间层，该中间层可直接访问目标IP和端口
			switch domainServer.targetRequest.Protocol { //对代理协议进行处理
			case Protocol_HTTP: //中间层代理或最内层代理，将请求进行直接转发，转发到目标IP和端口
				webRemoteStr = "http://" + domainServer.targetRequest.TargetIP + ":" + domainServer.targetRequest.TargetPort
			case Protocol_HTTPS:
			//todo 处理https请求
			case Protocol_WS:
				//todo 处理websocket

			default:
				return nil
			}
		}
	default:
		return nil

	}
	webremote, err = url.Parse(webRemoteStr)
	if err != nil {
		return err
	}
	domainServer.httpProxy = httputil.NewSingleHostReverseProxy(webremote)
	domainServer.httpsProxy = httputil.NewSingleHostReverseProxy(webremote)
	domainServers["a"] = domainServer

	return err
}
func ServerSetupWithOps(serverOpts ServerOptions) error {
	var (
		err error
	)

	httpServerListener, err = net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(serverOpts.HTTPPort))
	if err != nil {
		return err
	}
	//todo 需要在报错时，对上一步创建的httpserverproxy 进行资源释放
	httpsServerListener, err = net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(serverOpts.HTTPSPort))
	if err != nil {
		return err
	}
	sshServerListener, err = net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(serverOpts.SSHPort))
	if err != nil {
		return err
	}

	httpServerMuxRouter := mux.NewRouter()
	httpServerMuxRouter.PathPrefix("/").HandlerFunc(RouterHandler)
	//httpServerMuxRouter.Use() // tod 可以在此处进行中间件处理
	httpServer := &http.Server{
		Handler:      httpServerMuxRouter,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//	sshServer=& //todo  解决ssh代理
	go func() {
		if errStream <- httpServer.Serve(httpServerListener.(*net.TCPListener)); err != nil {
			log.Println(err)
			errStream <- err
		}
	}()

	go func() {
		httpServer.TLSConfig = nil // todo
		if err = httpServer.ServeTLS(httpsServerListener.(*net.TCPListener), "", ""); err != nil {
			log.Println(err)
			errStream <- err
		}
	}()
	log.Printf("proxy server staring ....,listen up at %s:%d\n", "0.0.0.0", serverOpts.HTTPPort)
	closeSignal := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(closeSignal, os.Interrupt)

	// Block until we receive our signal.
	<-closeSignal

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	httpServer.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("proxy server shutting down....")
	//os.Exit(0)

	return nil
}

//固定端口，通过泛域名形式进行路由转发
func RouterHandler(w http.ResponseWriter, r *http.Request) {
	var (
		server *PanDomainServer
	)

	server = domainServers["a"]
	/**
	//todo 域名处理，比如需要对域名进行校验，

	*/
	//reqDomain:=r.Host
	if r.Header.Get("Upgrade") == "websocket" {
		//todo 处理websocket请求

		//host := strings.SplitN(r.Header.Get("Origin"), "://", 2)
		//		addr := host[len(host)-1]
		//		r.Header.Set("Host", addr)
		server.webSocketProxy.ServeHTTP(w, r)
		return
	}
	if AuthAndFilterMiddleware(server.targetRequest, w, r) {
		switch server.targetRequest.Protocol {
		case Protocol_HTTP:
			server.httpProxy.ServeHTTP(w, r)
		case Protocol_HTTPS:
			server.httpsProxy.ServeHTTP(w, r)
		default:
			return
		}
	}
}

/**
代理服务中间认证和请求过滤层
*/
func AuthAndFilterMiddleware(proxyReq *TargetRequest) bool {
	if serverOpts.ServerLayer == User_Layer_Proxy { //外层，用户层代理，需要做权限验证
		switch proxyReq.SecuryBy {
		case SecureBy_None:
			return true
		case SecureBy_User: //todo 获取请求token，并进行token验证
			/*	userToken:= r.Header.Get(TokenName)
				if userToken==""{
					return false
				}*/
			return true
		default:
			return true
		}
	}
	//todo 新增过滤处理 filter 可以采用mux router中所具有的中间件处理过滤模式进行处理，只需要编写mux 中间件即可
	//参考 https://stackoverflow.com/questions/26204485/gorilla-mux-custom-middleware


	return true
}

/**
proxy server clean up
//todo 需要完善逻辑
*/
func ServerCleanup() error {
	var (
		err error
	)
	if httpServerListener != nil {
		err = httpServerListener.Close()

	}
	if httpsServerListener != nil {
		err = httpsServerListener.Close()
	}
	if sshServerListener != nil {
		err = sshServerListener.Close()

	}
	return err
}

//创建代理请求实例
func InitEntryInst(entry TargetRequest) error {

	var (
		err error
	//	entryInst EntryInstance
	)

	return err
}
