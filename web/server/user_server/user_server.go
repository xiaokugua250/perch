package main

import (
	"net/http"
	admin "perch/api/user_api"
	"perch/web/service"
)

func main() {

	serverRouter := []service.WebRouter{
		{RouterPath: "/sign_in", RouterHandlerFunc: admin.AuthUserSignInHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/sign_up", RouterHandlerFunc: admin.AuthUserSignUpHandler, RouterMethod: http.MethodPost, RouterDescription: "用户注册"},
		{RouterPath: "/info", RouterHandlerFunc: admin.AuthUserInfoHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/logout", RouterHandlerFunc: admin.PlatLogoutHandler, RouterMethod: http.MethodPost},

		{RouterPath: "/users", RouterHandlerFunc: admin.AuthUsersGetHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/user/{id}", RouterHandlerFunc: admin.PlatAuthUserUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/user/{id}", RouterHandlerFunc: admin.PlatAuthUserDeleteHandler, RouterMethod: http.MethodDelete},
		{RouterPath: "/user/{id}", RouterHandlerFunc: admin.PlatSpecAuthUserGetHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/roles", RouterHandlerFunc: admin.AuthRolesGetHandler, RouterMethod: http.MethodGet,RouterDescription: "获取用户角色"},
		{RouterPath: "/role/{id}", RouterHandlerFunc: admin.PlatAuthRoleUpdateHandler, RouterMethod: http.MethodPatch},

		{RouterPath: "/role/{id}", RouterHandlerFunc: admin.PlatAuthRoleDeleteHandler, RouterMethod: http.MethodDelete},
		{RouterPath: "/role/{id}", RouterHandlerFunc: admin.PlatSpecAuthRoleGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/role", RouterHandlerFunc: admin.PlatAuthRoleCreateHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/role/permissions{id}", RouterHandlerFunc: admin.PlatAuthSpecRolePermissionsGetHandler, RouterMethod: http.MethodGet, RouterDescription: "get spec role permissions"},

		{RouterPath: "/permissions", RouterHandlerFunc: admin.PlatAuthPermissionsGetHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/permission/{id}", RouterHandlerFunc: admin.PlatAuthPermissionUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/permission/{id}", RouterHandlerFunc: admin.PlatAuthPermissionDeleteHandler, RouterMethod: http.MethodDelete},
		{RouterPath: "/permission/{id}", RouterHandlerFunc: admin.PlatSpecAuthPermissionGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/permission", RouterHandlerFunc: admin.PlatAuthPermissionCreateHandler, RouterMethod: http.MethodPost},
	}
	webServer := service.NewWebServerWithOptions("user-micro", service.WithMySQLDBOptions(""))
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)
	webServer.InitFuncs = InitFunc
	webServer.Start()

}
