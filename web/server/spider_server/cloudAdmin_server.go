package main

import (
	"net/http"
	"perch/api/spider_api"
	"perch/web/service"
)

func main() {

	serverRouter := []service.WebRouter{
		{RouterPath: "/spiders", RouterHandlerFunc: spider_api.CreateCollySpider, RouterMethod: http.MethodPost},
	}

	webServer := service.NewWebServerWithOptions("plat-spider", service.WithMySQLDBOptions(""))
	webServer.Router = serverRouter
	webServer.Start()
}
