package main

import (
	"net/http"
	admin "perch/api/resources_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{
		{RouterPath: "/blogs", RouterHandlerFunc: admin.GetResourcesBlogsHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/blog", RouterHandlerFunc: admin.CreateResourcesBlogsHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/blog/{id}", RouterHandlerFunc: admin.SpecGetResourcesBlogsHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/blog/{id}", RouterHandlerFunc: admin.UpdateSpecResourcesBlogsHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/blog/{id}", RouterHandlerFunc: admin.DeleteSpecResourcesBlogsHandler, RouterMethod: http.MethodDelete},
		//{RouterPath: "/user/register", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},
		//------------------------
		{RouterPath: "/categorys", RouterHandlerFunc: admin.GetResourcesCategorysHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/category", RouterHandlerFunc: admin.CreateResourcesCategorysHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/category/{id}", RouterHandlerFunc: admin.SpecGetResourcesCategorysHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/category/{id}", RouterHandlerFunc: admin.UpdateSpecResourcesCategorysHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/category/{id}", RouterHandlerFunc: admin.DeleteSpecResourcesCategorysHandler, RouterMethod: http.MethodDelete},
	}

	webServer := service.NewWebServerWithOptions("plat-resources", service.WithMySQLDBOptions(""))
	webServer.Router = serverRouter
	webServer.Start()
}
