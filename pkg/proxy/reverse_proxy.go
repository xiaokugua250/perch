/**
golang 反向代理
*/
package proxy

import (
	log "github.com/sirupsen/logrus"
	//"github.com/opentracing/opentracing-go/log"
	"net/http"
	"net/http/httputil"
	"net/url"
	_ "perch/pkg/log"
	"strings"
)

/**
针对单个URL的反向代理
*/
func NewSingleHostReverseProxty(target string) *httputil.ReverseProxy {

	url, err := url.Parse(target)
	if err != nil {
		log.Error(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.Director = func(req *http.Request) {
		req.URL.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = url.Host

	}
	return reverseProxy

}

/**
针对具有多种路由前缀的反向代理
@targets:
    proxyMap := make(map[string]*url.URL)
	proxyMap ["plat-resources"]=
		&url.URL{
			Scheme:"http",
			Host:"127.0.0.1:8082",
	}
		proxyMap["plat-admin"]=&url.URL{
			Scheme:"http",
			Host:"127.0.0.1:8081",
		}
*/
func NewMultipleHostsReverseProxy(targets map[string]*url.URL) *httputil.ReverseProxy {

	if targets == nil {
		log.Error("router proxy target is empty")
		return nil
	}

	director := func(req *http.Request) {
		//println("CALLING DIRECTOR")
		prefix := strings.Split(req.URL.Path, "/")[1] //获取第一个URL参数来判断要路由的目标服务
		if prefix == "" {
			log.Errorf("router prefix is emtpy,value is:%s", prefix)
			return
		} else {
			target := targets[prefix]
			if target != nil {
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host

				//				req.URL.Path = target.String()+strings.Split(req.URL.Path,prefix)[1:][0]
				req.URL.Path = target.String() + strings.TrimPrefix(req.URL.Path, "/"+prefix)
				// 拼接转发URL demo :127.0.0.1:9090/plat-resources/resources/articles -->127.0.0.1:8082/resources/articles

			}
		}
		//	http://127.0.0.1:8082
	}
	return &httputil.ReverseProxy{
		Director: director,
	}
}
