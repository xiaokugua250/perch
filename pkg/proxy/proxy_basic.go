package proxy

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/websocket"
	"github.com/koding/websocketproxy"
	"golang.org/x/net/http2"
	"github.com/gorilla/mux"
)

const (
	Middle_Layer_Proxy = iota + 1 //中间层
	User_Layer_Proxy              //外层，用户层

)
const (
	Securit_None  = iota + 1 // 不需要验证用户
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

type ProxyServerConf struct{

}

type ProxyEntry struct {
	ID            string `json:"id"`
	UserName      string //用户名称
	EntryURL      string //代理URL
	ProtocolLayer string //代理所在层次
	Protocol      int    // 代理协议
	Domain        string //外部域名
	CertFile      []byte //ssl 证书
	KeyFile       []byte //ssl 证书
	TargrtIP      string
	TargetPort    string
	//TargetProtocl string
	SecuryBy int // 安全配置

}

// todo  https://github.com/golang/go/issues/26479
type ProxyServer struct {
	ProxyEntry *ProxyEntry
	httpProxy   *httputil.ReverseProxy
	httpsProxy *httputil.ReverseProxy
	webSocketProxy *websocketproxy.WebsocketProxy
}


// 启动代理服务

func ProxyServerSetup()(error){
	var (
		err error
	)

	defer func ListenerErrorHanlder(err *error){
		if err != nil{
		
			log.Printf("error in creating proxyserver is %s",err.Error())
		}
	}(&err)
	httpServerListener ,err = net.Listen("tcp",  "0.0.0.0:80")
	if err != nil {
		return err
	}
	//todo 需要在报错时，对上一步创建的httpserverproxy 进行资源释放
	httpsServerListener,err= net.Listen("tcp", "0.0.0.0:443")
	if err != nil {
		return err
	}
	sshServerListener,err = net.Listen("tcp", "0.0.0.0:22")
	if err != nil {
		return err
	}

	httpServerMuxRouter := mux.NewRouter()
	httpServerMuxRouter.HandlerFunc("/",ProxyRouterHandler)
	httpServer=&http.Server{
		Handler: 	httpServerMuxRouter ,
		WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
	}
//	sshServer=& //todo  解决ssh代理

	go func(){
		if err := httpServer.ListenAndServe(); err != nil {
            log.Println(err)
        }
	}()
	
	go func(){
		
		
		 
	}()

 	 closeSignal := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(closeSignal, os.Interrupt)

    // Block until we receive our signal.
    <-closeSignal

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    httpServer.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("shutting down")
    os.Exit(0)
}


func ProxyRouterHandler(){

}