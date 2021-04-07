/**
正向代理
*/
package wallproxy

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"perch/pkg/general/viperconf"
	"strconv"
	"syscall"
	"time"
)

var (
	serverIP    string
	serverPort  string
	proxyserver *proxyServer
)

type proxyServer struct {
	serverIP    string
	serverPort  string
	sslKeyfile  string
	sslCertfile string
}

func (proxy *proxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect { //代理模式

	} else {
		handleHTTP(w, r)

	}
}

func init() {
	proxyserver = &proxyServer{
		serverIP:    "0.0.0.0",
		serverPort:  "",
		sslKeyfile:  "",
		sslCertfile: "",
	}
}

func setup() {
	//webserver.Init()
	httpAddr := proxyserver.serverIP + ":" + proxyserver.serverPort

	//templ := `{{ .Title "Banner" "" 4 }}`
	banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString(fmt.Sprintf("{{ .Title \" %s \" \"\" 4 }}", webserver.Name)))
	fmt.Println()
	log.Println(" proxy service starting...")
	log.Println("service listening on：http://" + httpAddr)
	// 设置和启动服务
	server := &http.Server{
		Addr:    httpAddr,
		Handler: proxyserver,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.ListenAndServe()
		log.Println(" proxy server shutting down....")
	}()
	log.Println("proxy server start successfully....")

	// 捕获退出信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Println(err)
			}

			return
		case <-signalChan:
			server.Shutdown(ctx)
		}
	}
}

/**
代理模式下的数据处理方法
request.method==http.connect
*/
func handleTunneling(w http.ResponseWriter, req *http.Request) {

}

/**
直接访问模式下的http处理方式
指 GET\POST\PATCH\HEAD等请求时的处理逻辑
*/
func handleHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	handleHTTPHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

}

/**
拷贝复制http 请求头信息
*/
func handleHTTPHeader(dst, src http.Header) {
	for k, value := range src {
		for _, v := range value {
			dst.Add(k, v)
		}
	}
}

func clean() {

}
