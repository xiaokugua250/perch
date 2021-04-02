package main

import (
	"net/http"
	cloud "perch/api/k8scloud_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{
		{RouterPath: "/resources/namespaces", RouterHandlerFunc: cloud.CloudNameSpacesResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/resources/nodes", RouterHandlerFunc: cloud.CloudNodeResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/resources", RouterHandlerFunc: cloud.CloudResoucesHandler, RouterMethod: http.MethodGet},
		//----
		{RouterPath: "/construct/resource", RouterHandlerFunc: cloud.CloudResourceFileHandler, RouterMethod: http.MethodPost},
	}

	webServer := service.NewWebServerWithOptions("notify-micro", service.WithMySQLDBOptions(""), service.WithKubernetesOptions(""))
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)
	//InitFunc["database"]=database.InitMySQLDBWithConig
	webServer.InitFuncs = InitFunc
	webServer.Start()

}
