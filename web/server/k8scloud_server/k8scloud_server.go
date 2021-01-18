package main

import (
	"net/http"
	cloud "perch/api/k8scloud_api"
	"perch/web/service"
)

func main() {
	serverRouter := []service.WebRouter{

		{RouterPath: "/clusters", RouterHandlerFunc: cloud.CloudClustersHandler, RouterMethod: http.MethodGet},

		{RouterPath: "/namespaces", RouterHandlerFunc: cloud.CloudNameSpacesResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/nodes", RouterHandlerFunc: cloud.CloudNodeResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/workload/{cluster}", RouterHandlerFunc: cloud.CloudWorkloadHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/workload/{cluster}/{namespaces}/{type}", RouterHandlerFunc: cloud.CloudWorkloadResourcesHandler, RouterMethod: http.MethodGet},


		/*
		{RouterPath: "/service", RouterHandlerFunc: cloud.CloudServiceResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/configmap", RouterHandlerFunc: cloud.CloudConfigMapResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/serviceaccount", RouterHandlerFunc: cloud.CloudServiceAccountResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/pod", RouterHandlerFunc: cloud.CloudPODResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/job", RouterHandlerFunc: cloud.CloudJOBResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/batchjob", RouterHandlerFunc: cloud.CloudBatchJOBResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/deployment", RouterHandlerFunc: cloud.CloudDeploymentResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/daemonset", RouterHandlerFunc: cloud.CloudDaemonSetResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/replicaset", RouterHandlerFunc: cloud.CloudReplicasetResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/statefulset", RouterHandlerFunc: cloud.CloudStatefuleSetResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/pv", RouterHandlerFunc: cloud.CloudPVResoucesHandler, RouterMethod: http.MethodGet},
		{RouterPath: "/pvc", RouterHandlerFunc: cloud.CloudPVCResoucesHandler, RouterMethod: http.MethodGet},
		*/
		{RouterPath: "/resources", RouterHandlerFunc: cloud.CloudResoucesHandler, RouterMethod: http.MethodGet},
		//----
		{RouterPath: "/construct/resource", RouterHandlerFunc: cloud.CloudResourceFileHandler, RouterMethod: http.MethodPost},
	}

	webServer := service.NewWebServerWithOptions("k8scloud-micro", service.WithMySQLDBOptions(""), service.WithKubernetesOptions(""))
	webServer.Router = serverRouter

	InitFunc := make(map[string]func(config interface{}) error)

	webServer.InitFuncs = InitFunc
	webServer.Start()

}
