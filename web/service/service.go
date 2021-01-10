package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"perch/pkg/general/viperconf"
	"strconv"
	"syscall"
	"time"
)

type WebRouter struct {
	RouterPath        string                                         `json:"router_path"`
	RouterHandlerFunc func(w http.ResponseWriter, req *http.Request) `json:"router_handler_func"`
	RouterMethod      string                                         `json:"router_method"`
	RouterPathPrefiex bool                                           `json:"router_path_prefiex" ` //前置路由匹配
	RouterDescription string                                         `json:"router_description"`
}

type WebServer struct {
	Name   string      `json:"name"`
	Router []WebRouter `json:"router"`
	//InitFunc           []func() error `json:"init_func"`
	//InitFuncConfigMaps map[string]interface{}                    `json:"init_func_config_maps"` //对应初始函数名和所需要的配置参数路径
	InitFuncs map[string]func(config interface{}) error `json:"init_funcs"`
	CleanFunc []func() error                            `json:"clean_func"`
}

func (webserver *WebServer) GenRouter() *mux.Router {
	router := mux.NewRouter()
	/*	router.Use(middleware.CROSMiddleware)
		router.Use(middleware.MetricMiddleWare)
		router.Use(middleware.LoggingMiddleware)*/
	for _, r := range webserver.Router {
		if r.RouterPathPrefiex {
			router.Methods(r.RouterMethod).PathPrefix(r.RouterPath).Path(r.RouterPath).HandlerFunc(r.RouterHandlerFunc)
		} else {
			router.Methods(r.RouterMethod).Path(r.RouterPath).HandlerFunc(r.RouterHandlerFunc)
		}
	}
	return router
}
func (webserver *WebServer) Init() {
	err := viperconf.InitGenWebConfig(webserver.Name)
	if err != nil {
		log.Fatalln(err)
	}

}

func (webserver *WebServer) Start() {

	//webserver.Init()
	httpAddr := viperconf.WebServiceConfig.WebConfig.ServerIP + ":" + strconv.Itoa(viperconf.WebServiceConfig.WebConfig.ServerPort)

	log.Println(webserver.Name + " service starting...")
	log.Println("service listening on：http://" + httpAddr)
	// 设置和启动服务
	server := &http.Server{
		Addr: httpAddr,
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type", "Cache-Control", "x-token", "ETag", "TIMEOUT", "DEADLINE", "content-range", "application/json"}),
		)(webserver.GenRouter()),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.ListenAndServe()
		log.Println(webserver.Name + " shutting down....")
	}()
	log.Println(webserver.Name + " start successfully....")

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
			// 执行额外的清理操作
			for _, clean := range webserver.CleanFunc {
				clean()
			}
			return
		case <-signalChan:
			server.Shutdown(ctx)
		}
	}
}
