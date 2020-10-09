package main

import (
	"net/http"
	admin "perch/api/dataplat_api"
	database "perch/database/mysql"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "plat-resources",

		Router: []service.WebRouter{
			{RouterPath: "/resources/articles", RouterHandlerFunc: admin.PlatDataResourcesHandler, RouterMethod: http.MethodGet},

			//{RouterPath: "/user/register", RouterHandlerFunc: admin.PlatLoginGenTokenHandler, RouterMethod: http.MethodPost},


		},
		InitFunc: []func() error{
			database.InitMySQLDB,
		},
	}.WebServiceStart()

}
