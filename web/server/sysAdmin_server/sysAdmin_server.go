package main

import (
	"net/http"
	sysadmin "perch/api/system_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{
		{RouterPath: "/sys/memadvanced", RouterHandlerFunc: sysadmin.SysMemInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/sys/hostadvanced", RouterHandlerFunc: sysadmin.SysHostInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/sys/cpuadvanced", RouterHandlerFunc: sysadmin.SysCpuInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/sys/dockeradvanced", RouterHandlerFunc: sysadmin.SysDockerInfoHandler, RouterMethod: http.MethodGet},

		//	{RouterPath: "/sys/diskadvanced", RouterHandlerFunc: sysadmin.SysDiskInfoHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/sys/loadadvanced", RouterHandlerFunc: sysadmin.SysLoadInfoHandler, RouterMethod: http.MethodGet},

		//{RouterPath: "/sys/netadvanced", RouterHandlerFunc: sysadmin.SysNetInfoHandler, RouterMethod: http.MethodGet}, //todo 方法需完善

		{RouterPath: "/sys/processadvanced", RouterHandlerFunc: sysadmin.SysProcessInfoHandler, RouterMethod: http.MethodGet}, //todo 方法需完善

	}
	webServer := service.NewWebServerWithOptions("sys-admin", service.WithDatabaseOptions(""))
	webServer.Router = serverRouter
	webServer.Start()
}
