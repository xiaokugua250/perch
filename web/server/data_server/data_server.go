package main

import (
	"net/http"
	"perch/api"
	admin "perch/api/user_api"
	"perch/web/service"
)

func main() {

	serverRouter := []service.WebRouter{
		{RouterPath: "/version", RouterHandlerFunc: api.ServiceVersionandler, RouterMethod: http.MethodGet},
		{RouterPath: "/health", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},
		
		{RouterPath: "/fake/users", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/emails", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/creditCard", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/ip", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/loc", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/times", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},


	}
	webServer := service.NewWebServerWithOptions("plat-data",)
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)
	//InitFunc["database"]=database.InitMySQLDBWithConig
	webServer.InitFuncs = InitFunc
	webServer.Start()

}
