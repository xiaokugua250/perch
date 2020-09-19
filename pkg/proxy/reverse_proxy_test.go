
/**
golang 反向代理
 */
package proxy

import (
	"log"
	"net/http"

	"net/url"
	"testing"
)

func TestNewMultipleHostsReverseProxy(t *testing.T) {
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
	reverProxy :=  NewMultipleHostsReverseProxy(proxyMap)

	//reverProxy.ServeHTTP()

	log.Fatal(http.ListenAndServe(":9090",reverProxy))


}

