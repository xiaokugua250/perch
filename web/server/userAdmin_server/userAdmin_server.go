package main

import (
	"net/http"
	"perch/api"
	admin "perch/api/user_api"
	database "perch/database/mysql"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "plat-admin",
		Router: []service.WebRouter{
			{RouterPath: "/version", RouterHandlerFunc: api.ServiceVersionandler, RouterMethod: http.MethodGet},
			{RouterPath: "/health", RouterHandlerFunc: api.ServiceHealthHandler, RouterMethod: http.MethodGet},

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
			//{RouterPath: "/user/register", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},

			{RouterPath: "/auth-rbac/roles", RouterHandlerFunc: admin.PlatAuthRolesGetHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/auth-rbac/role/{id}", RouterHandlerFunc: admin.PlatAuthRoleUpdateHandler, RouterMethod: http.MethodPatch},
			{RouterPath: "/auth-rbac/role/{id}", RouterHandlerFunc: admin.PlatAuthRoleDeleteHandler, RouterMethod: http.MethodDelete},
			{RouterPath: "/auth-rbac/role/{id}", RouterHandlerFunc: admin.PlatSpecAuthRoleGetHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/auth-rbac/role", RouterHandlerFunc: admin.PlatAuthRoleCreateHandler, RouterMethod: http.MethodPost},
			{RouterPath: "/auth-rbac/role/permissions{id}", RouterHandlerFunc: admin.PlatAuthSpecRolePermissionsGetHandler, RouterMethod: http.MethodGet, RouterInfo: "get spec role permissions"},
			{RouterPath: "/auth-rbac/permissions", RouterHandlerFunc: admin.PlatAuthPermissionsGetHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/auth-rbac/permission/{id}", RouterHandlerFunc: admin.PlatAuthPermissionUpdateHandler, RouterMethod: http.MethodPatch},
			{RouterPath: "/auth-rbac/permission/{id}", RouterHandlerFunc: admin.PlatAuthPermissionDeleteHandler, RouterMethod: http.MethodDelete},
			{RouterPath: "/auth-rbac/permission/{id}", RouterHandlerFunc: admin.PlatSpecAuthPermissionGetHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/auth-rbac/permission", RouterHandlerFunc: admin.PlatAuthPermissionCreateHandler, RouterMethod: http.MethodPost},
		},
		InitFunc: []func() error{
			database.InitMySQLDB,
		},
	}.WebServiceStart()

}
