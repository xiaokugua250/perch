package main

import (
	"net/http"
	sysadmin "perch/api/system_api"
	"perch/web/service"
)

func main() {
	//
	serverRouter := []service.WebRouter{
		{RouterPath: "/basicinfo", RouterHandlerFunc: sysadmin.SysBasicInfoHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/memadvanced", RouterHandlerFunc: sysadmin.SysMemInfoHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/hostadvanced", RouterHandlerFunc: sysadmin.SysHostInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/cpuadvanced", RouterHandlerFunc: sysadmin.SysCpuInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/dockeradvanced", RouterHandlerFunc: sysadmin.SysDockerInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/diskadvanced", RouterHandlerFunc: sysadmin.SysDiskInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/loadadvanced", RouterHandlerFunc: sysadmin.SysLoadInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/netadvanced", RouterHandlerFunc: sysadmin.SysNetInfoHandler, RouterMethod: http.MethodGet}, //todo 方法需完善

		{RouterPath: "/processadvanced", RouterHandlerFunc: sysadmin.SysProcessInfoHandler, RouterMethod: http.MethodGet}, //todo 方法需完善

	}
	webServer := service.NewWebServerWithOptions("sys-micro")
	webServer.Router = serverRouter
	webServer.Start()
}
