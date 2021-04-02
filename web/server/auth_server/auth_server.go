package main

/**
网站基本信息
*/
import (
	"net/http"
	"perch/api/auth_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{

		{RouterPath: "/casbin/auth", RouterHandlerFunc: auth_api.CasbinAuthCreateHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/casbin/auth/{id}", RouterHandlerFunc: auth_api.CasbinAuthSpcGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/casbin/auths", RouterHandlerFunc: auth_api.CasbinAuthGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/casbin/auths/{id}", RouterHandlerFunc: auth_api.CasbinAuthPatchHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/casbin/auths/{id}", RouterHandlerFunc: auth_api.CasbinAuthDeleteHandler, RouterMethod: http.MethodDelete},
	}

	webServer := service.NewWebServerWithOptions("auth-micro", service.WithMySQLDBOptions(""))
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)

	webServer.InitFuncs = InitFunc
	webServer.Start()

}
