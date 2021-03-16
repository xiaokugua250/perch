package proxy

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	TokenName = "Secret-Token"
)

type ServerOptions struct {
	HTTPPort int
	HTTPSPort int
	SSHPort int
	MiddleLayer string // 中间代理层
	SSLCertFile string
	SSLKeyFile string
}

type ProxyEntry struct {
	ID         string `json:"id"`
	UserName   string //用户名称
	EntryURL   string //代理URL
	Layer      int    //代理所在层次
	Protocol   int    // 代理协议
	Domain     string //外部域名
	SSLCertFile   []byte //ssl 证书
	SSLKeyFile    []byte //ssl 证书
	TargrtIP   string
	TargetPort string
	//TargetProtocl string
	SecuryBy int // 安全配置

}

// todo  https://github.com/golang/go/issues/26479
//泛域名代理服务器
type PanDomainServer struct {
	ProxyEntry     *ProxyEntry
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
	serverOptions ServerOptions //
)
// 启动代理服务

func ServerSetup() error {
	var (
		err error
	)
	/*	defer func ListenerErrorHanlder(err *error){
		if err != nil{
			log.Printf("error in creating proxyserver is %s",err.Error())
		}
	}(&err)*/
	httpServerListener, err = net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(serverOptions.HTTPPort))
	if err != nil {
		return err
	}
	//todo 需要在报错时，对上一步创建的httpserverproxy 进行资源释放
	httpsServerListener, err = net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(serverOptions.HTTPSPort))
	if err != nil {
		return err
	}
	sshServerListener, err = net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(serverOptions.SSHPort))
	if err != nil {
		return err
	}

	httpServerMuxRouter := mux.NewRouter()
	httpServerMuxRouter.HandleFunc("/", RouterHandler)
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
	/*closeSignal := make(chan os.Signal, 1)
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
	log.Println("proxy server shutting down....")*/
	//os.Exit(0)
	return nil
}


//固定端口，通过泛域名形式进行路由转发
func RouterHandler(w http.ResponseWriter, r *http.Request) {
	var (
		server *PanDomainServer
	)
	/**
	//todo 域名处理，比如需要对域名进行校验，

	 */
	//reqDomain:=r.Host

	if r.Header.Get("Upgrade")=="websocket"{
		//todo 处理websocket请求

		//host := strings.SplitN(r.Header.Get("Origin"), "://", 2)
		//		addr := host[len(host)-1]
		//		r.Header.Set("Host", addr)
		server.webSocketProxy.ServeHTTP(w,r)
		return
	}
	if AuthAndFilterLayer(server.ProxyEntry,w,r){
		switch server.ProxyEntry.Protocol {
		case Protocol_HTTP:
			server.httpProxy.ServeHTTP(w,r)
		case Protocol_HTTPS:
			server.httpsProxy.ServeHTTP(w,r)
		default:
			return
		}
	}
}

/**
代理服务中间认证和请求过滤层
*/
func AuthAndFilterLayer(proxyReq *ProxyEntry, w http.ResponseWriter, r *http.Request) bool {

	if proxyReq.Layer == User_Layer_Proxy { //外层，用户层代理，需要做权限验证
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

	return true
}


/**
proxy server clean up
//todo 需要完善逻辑
 */
func ServerCleanup()error{
	var (
		err error
	)
	if httpServerListener!= nil{
		err = httpServerListener.Close()

	}
	if httpsServerListener !=nil{
		err= httpsServerListener.Close()
	}
	if sshServerListener!= nil{
		err= sshServerListener.Close()

	}
	return err
}
