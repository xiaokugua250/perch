package service

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"perch/web/middleware"
	"syscall"
	"time"
)

type WebRouter struct {
	RouterPath        string                                         `json:"router_path"`
	RouterHandlerFunc func(w http.ResponseWriter, req *http.Request) `json:"router_handler_func"`
	RouterMethod      string                                         `json:"router_method"`
	RouterPathPrefiex bool                                           `json:"router_path_prefiex" ` //前置路由匹配
	RouterInfo        string                                         `json:"router_info"`
}

type WebService struct {
	Name      string
	Router    []WebRouter
	InitFunc  []func() error
	CleanFunc []func() error
}

/**
通用初始化方法
如:初始化数据库、配置文件解析等
*/
func GeneralInitFunc() error {
	return nil
}

func GeneralCleanFunc() error {
	return nil
}

func (webservice WebService) WebServiceInit() {
	for _, initFunc := range webservice.InitFunc {
		err := initFunc()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (webservice WebService) WebServiceGenRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.CROSMiddleware)
	router.Use(middleware.MetricMiddleWare)
	router.Use(middleware.LoggingMiddleware)
	for _, r := range webservice.Router {
		if r.RouterPathPrefiex {
			router.Methods(r.RouterMethod).PathPrefix(r.RouterPath).Path(r.RouterPath).HandlerFunc(r.RouterHandlerFunc)
		} else {
			router.Methods(r.RouterMethod).Path(r.RouterPath).HandlerFunc(r.RouterHandlerFunc)
		}
	}
	return router
}
func (webservice WebService) WebServiceStart() {
	webservice.WebServiceInit()
	httpAddr := "0.0.0.0:8081"

	log.Println(webservice.Name + " service starting...")
	log.Println("service listening on：http://" + httpAddr)
	// 设置和启动服务
	server := &http.Server{
		Addr:         httpAddr,
		Handler:      webservice.WebServiceGenRouter(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.ListenAndServe()
		log.Println(webservice.Name + "shutting down....")
	}()
	log.Println(webservice.Name + "start successfully....")

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
			for _, clean := range webservice.CleanFunc {
				clean()
			}

			return
		case <-signalChan:

			server.Shutdown(ctx)
		}
	}

}
