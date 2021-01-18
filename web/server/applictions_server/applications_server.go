package main

import (
	"net/http"
	app "perch/api/applications_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{
		{RouterPath: "/application", RouterHandlerFunc: app.ApplicationsCreateHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/applications", RouterHandlerFunc: app.ApplicationsGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/applications/{id}", RouterHandlerFunc: app.ApplicationsSpecGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/applications/{id}", RouterHandlerFunc: app.ApplicationsSpecUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/applications/{id}", RouterHandlerFunc: app.ApplicationsSpecDeleteHandler, RouterMethod: http.MethodDelete},

		{RouterPath: "/instances", RouterHandlerFunc: app.ApplicationsInstancesCreateHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/instances", RouterHandlerFunc: app.ApplicationsInstancesGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/instances/{id}", RouterHandlerFunc: app.ApplicationsInstancesSpecGetHandler, RouterMethod: http.MethodDelete},

		{RouterPath: "/instances/{id}", RouterHandlerFunc: app.ApplicationsInstancesSpecUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/instances/{id}", RouterHandlerFunc: app.ApplicationsInstancesSpecDeleteHandler, RouterMethod: http.MethodDelete},

	}

	webServer := service.NewWebServerWithOptions("application-micro", service.WithMySQLDBOptions(""), service.WithKubernetesOptions(""))
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)
	//InitFunc["database"]=database.InitMySQLDBWithConig
	webServer.InitFuncs = InitFunc
	webServer.Start()

}
