package main

import (
	"net/http"

	 api "perch/api/data_api"
	"perch/web/service"
)

func main() {

	serverRouter := []service.WebRouter{

		{RouterPath: "/fake/users", RouterHandlerFunc: api.GetFakeUsersHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/emails", RouterHandlerFunc: api.GetFakeEmailsHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/creditCard", RouterHandlerFunc: api.GetFakeCrediCardHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/ip", RouterHandlerFunc: api.GetFakeIPHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/loc", RouterHandlerFunc: api.GetFakeLocHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/fake/times", RouterHandlerFunc: api.GetFakeTimesHandler, RouterMethod: http.MethodGet},


	}
	webServer := service.NewWebServerWithOptions("plat-data",)
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)
	//InitFunc["database"]=database.InitMySQLDBWithConig
	webServer.InitFuncs = InitFunc
	webServer.Start()

}
