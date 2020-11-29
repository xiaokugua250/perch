package main

import (
	"net/http"
	admin "perch/api/resources_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{
		{RouterPath: "/docs", RouterHandlerFunc: admin.GetResourcesDocsHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/docs", RouterHandlerFunc: admin.CreateResourcesDocsHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/docs/{id}", RouterHandlerFunc: admin.SpecGetResourcesDocsHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/docs/{id}", RouterHandlerFunc: admin.UpdateSpecResourcesDocsHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/docs/{id}", RouterHandlerFunc: admin.DeleteSpecResourcesDocsHandler, RouterMethod: http.MethodDelete},
		//{RouterPath: "/user/register", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},
		//------------------------
		{RouterPath: "/categorys", RouterHandlerFunc: admin.GetResourcesCategorysHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/category", RouterHandlerFunc: admin.CreateResourcesCategorysHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/category/{id}", RouterHandlerFunc: admin.SpecGetResourcesCategorysHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/category/{id}", RouterHandlerFunc: admin.UpdateSpecResourcesCategorysHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/category/{id}", RouterHandlerFunc: admin.DeleteSpecResourcesCategorysHandler, RouterMethod: http.MethodDelete},
	}

	webServer := service.NewWebServerWithOptions("plat-resources", service.WithDatabaseOptions(""))
	webServer.Router = serverRouter
	webServer.Start()
}
