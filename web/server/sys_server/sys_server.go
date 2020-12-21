package main

import (
	"net/http"
	sysadmin "perch/api/system_api"
	"perch/web/service"
)

func main() {
	//
	serverRouter := []service.WebRouter{
		{RouterPath: "/monitor/basicinfo", RouterHandlerFunc: sysadmin.SysBasicInfoHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/monitor/memadvanced", RouterHandlerFunc: sysadmin.SysMemInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/monitor/hostadvanced", RouterHandlerFunc: sysadmin.SysHostInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/monitor/cpuadvanced", RouterHandlerFunc: sysadmin.SysCpuInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/monitor/dockeradvanced", RouterHandlerFunc: sysadmin.SysDockerInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/monitor/diskadvanced", RouterHandlerFunc: sysadmin.SysDiskInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/monitor/loadadvanced", RouterHandlerFunc: sysadmin.SysLoadInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/monitor/netadvanced", RouterHandlerFunc: sysadmin.SysNetInfoHandler, RouterMethod: http.MethodGet}, //todo 方法需完善

		{RouterPath: "/monitor/processadvanced", RouterHandlerFunc: sysadmin.SysProcessInfoHandler, RouterMethod: http.MethodGet}, //todo 方法需完善

	}
	webServer := service.NewWebServerWithOptions("sys-admin")
	webServer.Router = serverRouter
	webServer.Start()
}
