/**
正向代理
*/
package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
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
	proxyserver = &proxyServer{}

	configfile string
)

const (
	protocol_http  = "http"
	protocol_https = "https"
)

func LoadServerConfigs(configfile string) error {
	var (
		err error
	)
	conifgyaml, err := ioutil.ReadFile(configfile)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(conifgyaml, &proxyserver); err != nil {
		return err
	}
	return nil
}

type proxyServer struct {
	ServerIP    string `yaml:"ip"`
	ServerPort  string `yaml:"port"`
	Protocol    string `yaml:"protocol"`
	SSLKeyfile  string `yaml:"keyfile"`
	SSLCertfile string `yaml:"certfile"`
}

func (proxy *proxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodConnect { //代理模式
		/*	w.Header().Set("Content-Type", "text/html; charset=ascii")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
		*/
		handleWithProxy(w, r)
		log.Printf("handle http  request %s with proxy...\n", r.RequestURI)
	} else {
		HandleWithoutProxy(w, r)
		log.Printf("handle http request %s without proxy...\n", r.RequestURI)
	}
}

func setup() {
	var (
		//configfile string
		err error
	)
	if runtime.GOOS == "windows" {
		//todo
	} else if runtime.GOOS == "linux" {
		//todo
		if configfile != "" {
			configfile = "/etc/wallproxy/configs.yaml"
		}

	} else {
		log.Fatalf("current os %s not support...", runtime.GOOS)
	}
	err = LoadServerConfigs(configfile)
	if err != nil {
		log.Fatalf("load configfile failed,error is %s", err.Error())
	}
	proxyserver = &proxyServer{
		ServerIP:    proxyserver.ServerIP,
		ServerPort:  proxyserver.ServerPort,
		Protocol:    proxyserver.Protocol,
		SSLKeyfile:  proxyserver.SSLKeyfile,
		SSLCertfile: proxyserver.SSLCertfile,
	}
}

func main() {
	//webserver.Init()

	app := &cli.App{
		Name:     "wallproxy",
		Usage:    "cross the wall...",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "liangdu",
				Email: "liangdu1992@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "configs",
				//Value:       "/etc/wallproxy/configs.yaml",
				Usage:       "server proxy config file",
				Destination: &configfile,
			},
			&cli.StringFlag{
				Name:        "ip",
				Value:       "0.0.0.0",
				Usage:       "server proxy server ip",
				Destination: &proxyserver.ServerIP,
			},
			&cli.StringFlag{
				Name:        "port",
				Value:       "2578",
				Usage:       "server proxy server port",
				Destination: &proxyserver.ServerPort,
			},
			&cli.StringFlag{
				Name:        "protocol",
				Value:       protocol_http,
				Usage:       "server proxy server protocol",
				Destination: &proxyserver.Protocol,
			},
			&cli.StringFlag{
				Name:        "sslkeys",
				Value:       "",
				Usage:       "server proxy ssl keys",
				Destination: &proxyserver.SSLKeyfile,
			},
			&cli.StringFlag{
				Name:        "sslcerts",
				Value:       "",
				Usage:       "server proxy ssl certs",
				Destination: &proxyserver.SSLCertfile,
			},
		},
		Action: func(c *cli.Context) error {

			fmt.Println("enjoy freedom...")
			if configfile != "" {
				setup()
			}
			startUp()

			if c.NArg() <= 0 {
				return nil
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func startUp() {
	httpAddr := proxyserver.ServerIP + ":" + proxyserver.ServerPort

	//templ := `{{ .Title "Banner" "" 4 }}`
	banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString(fmt.Sprintf("{{ .Title \" %s \" \"\" 4 }}", "proxy_server")))
	fmt.Println()
	log.Println(" proxy service starting...")
	if proxyserver.Protocol == protocol_http {
		log.Println("service listening on:http://" + httpAddr)
	} else if proxyserver.Protocol == protocol_https {
		log.Println("service listening on: https://" + httpAddr)
	} else {
		log.Println("service listening on:", proxyserver.Protocol+"://"+httpAddr)
	}

	// 设置和启动服务
	server := &http.Server{
		Addr:    httpAddr,
		Handler: AuthMiddlerware(proxyserver),

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	errChan := make(chan error, 1)
	go func() {
		if proxyserver.Protocol == protocol_http {
			errChan <- server.ListenAndServe()
		} else {
			errChan <- server.ListenAndServeTLS(proxyserver.SSLCertfile, proxyserver.SSLKeyfile)
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

/**

认证中间件
*/
func AuthMiddlerware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo Do stuff here

		//log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)

	})
}
