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
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	serverIP    string
	serverPort  string
	proxyserver *proxyServer
	configs     ServerConfig
)

const (
	protocol_http = iota + 1
	protocol_https
)

type ServerConfig struct {
	serverIP    string `yaml:"ip"`
	serverPort  string `yaml:"ip"`
	protocol    int    `yaml:"ip"`
	sslKeyfile  string `yaml:"ip"`
	sslCertfile string `yaml:"ip"`
}

func LoadServerConfigs(configfile string) error {
	var (
		err error
	)
	conifgyaml, err := ioutil.ReadFile(configfile)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(conifgyaml, &configs); err != nil {
		return err
	}
	return nil
}

type proxyServer struct {
	serverIP    string
	serverPort  string
	protocol    int
	sslKeyfile  string
	sslCertfile string
}

func (proxy *proxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect { //代理模式
		handleWithProxy(w, r)
		log.Printf("handle http  request %s with proxy...\n", r.RequestURI)
	} else {
		HandleWithoutProxy(w, r)
		log.Printf("handle http request %s without proxy...\n", r.RequestURI)
	}
}

func init() {
	var (
		configfile string
		err        error
	)
	if runtime.GOOS == "windows" {
		//todo
	} else if runtime.GOOS == "linux" {
		//todo
		configfile = "/etc/wallproxy/server.yaml"
	} else {
		log.Fatalf("current os %s not support...", runtime.GOOS)
	}
	err = LoadServerConfigs(configfile)
	if err != nil {
		log.Fatalf("load configfile failed,error is %s", err.Error())
	}
	proxyserver = &proxyServer{
		serverIP:    configs.serverIP,
		serverPort:  configs.serverPort,
		protocol:    configs.protocol,
		sslKeyfile:  configs.sslKeyfile,
		sslCertfile: configs.sslCertfile,
	}
}

func Setup() {
	//webserver.Init()
	httpAddr := proxyserver.serverIP + ":" + proxyserver.serverPort

	//templ := `{{ .Title "Banner" "" 4 }}`
	banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString(fmt.Sprintf("{{ .Title \" %s \" \"\" 4 }}", "proxy_server")))
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
		if proxyserver.protocol == protocol_http {
			errChan <- server.ListenAndServe()
		} else {
			errChan <- server.ListenAndServeTLS(proxyserver.sslCertfile, proxyserver.sslKeyfile)
		}

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
func handleWithProxy(w http.ResponseWriter, req *http.Request) {
	destConn, err := net.DialTimeout("tcp", req.Host, time.Second*10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker) // hijeck 参考 https://stackoverflow.com/questions/27075478/when-to-use-hijack-in-golang
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

/**
直接访问模式下的http处理方式
指 GET\POST\PATCH\HEAD等请求时的处理逻辑
*/
func HandleWithoutProxy(w http.ResponseWriter, req *http.Request) {
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

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
func clean() {

}
