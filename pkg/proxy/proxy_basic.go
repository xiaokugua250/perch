package proxy

import (
	"net/http/httputil"
	"github.com/gorilla/websocket"
	"github.com/koding/websocketproxy"
	"golang.org/x/net/http2"
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
