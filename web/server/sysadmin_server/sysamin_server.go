package main

import (
	"net/http"
	sysadmin "perch/api/sysadmin_api"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "PEX_SYSINFO",

		Router: []service.WebRouter{
			{RouterPath: "/sys/mem", RouterHandlerFunc: sysadmin.SysMemInfoHandler, RouterMethod: http.MethodGet},

			{RouterPath: "/sys/hostadvanced", RouterHandlerFunc: sysadmin.SysHostInfoHandler, RouterMethod: http.MethodGet},
		},
		InitFunc: []func() error{
			//database.InitMySQLDB,
		},
	}.WebServiceStart()

}
