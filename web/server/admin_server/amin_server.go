package main

import (
	"net/http"
	admin "perch/api/admin_api"
	database "perch/database/mysql"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "PEX_ADMIN",

		Router: []service.WebRouter{
			{RouterPath: "/user/login", RouterHandlerFunc: admin.PlatLoginHandler, RouterMethod: http.MethodPost},
			{RouterPath: "/user/logout", RouterHandlerFunc: admin.PlatLogoutHandler, RouterMethod: http.MethodPost},
			{RouterPath: "/user/info", RouterHandlerFunc: admin.PlatUserInfoHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/user/admin", RouterHandlerFunc: admin.PlatAdminHandler, RouterMethod: http.MethodPost},
			{RouterPath: "/user/token", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},
			{RouterPath: "/auth-user/users", RouterHandlerFunc: admin. PlatAuthUsersGetHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/auth-user/user/{id}", RouterHandlerFunc: admin. PlatAuthUserUpdateHandler, RouterMethod: http.MethodPatch},
			{RouterPath: "/auth-user/user/{id}", RouterHandlerFunc: admin. PlatAuthUserDeleteHandler, RouterMethod: http.MethodDelete},
			{RouterPath: "/auth-user/user", RouterHandlerFunc: admin. PlatAuthUserCreateHandler, RouterMethod: http.MethodPost},
			//{RouterPath: "/user/register", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},


		},
		InitFunc: []func() error{
			database.InitMySQLDB,
		},
	}.WebServiceStart()

}
