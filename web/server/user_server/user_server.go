package main

import (
	"net/http"
	admin "perch/api/user_api"
	"perch/web/service"
)

func main() {

	serverRouter := []service.WebRouter{

		{RouterPath: "/user/login", RouterHandlerFunc: admin.PlatLoginHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/user/logout", RouterHandlerFunc: admin.PlatLogoutHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/user/info", RouterHandlerFunc: admin.PlatUserInfoHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/user/admin", RouterHandlerFunc: admin.PlatAdminHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/user/token", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/auth-user/users", RouterHandlerFunc: admin.PlatAuthUsersGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/auth-user/user/{id}", RouterHandlerFunc: admin.PlatAuthUserUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/auth-user/user/{id}", RouterHandlerFunc: admin.PlatAuthUserDeleteHandler, RouterMethod: http.MethodDelete},
		{RouterPath: "/auth-user/user/{id}", RouterHandlerFunc: admin.PlatSpecAuthUserGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/auth-user/user", RouterHandlerFunc: admin.PlatAuthUserCreateHandler, RouterMethod: http.MethodPost},

		{RouterPath: "/auth-rbac/roles", RouterHandlerFunc: admin.PlatAuthRolesGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/auth-rbac/role/{id}", RouterHandlerFunc: admin.PlatAuthRoleUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/auth-rbac/role/{id}", RouterHandlerFunc: admin.PlatAuthRoleDeleteHandler, RouterMethod: http.MethodDelete},
		{RouterPath: "/auth-rbac/role/{id}", RouterHandlerFunc: admin.PlatSpecAuthRoleGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/auth-rbac/role", RouterHandlerFunc: admin.PlatAuthRoleCreateHandler, RouterMethod: http.MethodPost},
		{RouterPath: "/auth-rbac/role/permissions{id}", RouterHandlerFunc: admin.PlatAuthSpecRolePermissionsGetHandler, RouterMethod: http.MethodGet, RouterDescription: "get spec role permissions"},
		{RouterPath: "/auth-rbac/permissions", RouterHandlerFunc: admin.PlatAuthPermissionsGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/auth-rbac/permission/{id}", RouterHandlerFunc: admin.PlatAuthPermissionUpdateHandler, RouterMethod: http.MethodPatch},
		{RouterPath: "/auth-rbac/permission/{id}", RouterHandlerFunc: admin.PlatAuthPermissionDeleteHandler, RouterMethod: http.MethodDelete},
		{RouterPath: "/auth-rbac/permission/{id}", RouterHandlerFunc: admin.PlatSpecAuthPermissionGetHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/auth-rbac/permission", RouterHandlerFunc: admin.PlatAuthPermissionCreateHandler, RouterMethod: http.MethodPost},
	}
	webServer := service.NewWebServerWithOptions("user-micro", service.WithMySQLDBOptions(""))
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)
	//InitFunc["database"]=database.InitMySQLDBWithConig
	webServer.InitFuncs = InitFunc
	webServer.Start()

}
