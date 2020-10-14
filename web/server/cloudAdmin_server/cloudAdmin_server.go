package main

import (
	"net/http"
	cloud "perch/api/cloud_api"
	database "perch/database/mysql"
	"perch/pkg/cluster/k8s"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "plat-cloud",

		Router: []service.WebRouter{
			{RouterPath: "/resources/namespaces", RouterHandlerFunc: cloud.CloudNameSpacesResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/nodes", RouterHandlerFunc: cloud.CloudNodeResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/service", RouterHandlerFunc: cloud.CloudServiceResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/configmap", RouterHandlerFunc: cloud.CloudConfigMapResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/serviceaccount", RouterHandlerFunc: cloud.CloudServiceAccountResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/pod", RouterHandlerFunc: cloud.CloudPODResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/job", RouterHandlerFunc: cloud.CloudJOBResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/batchjob", RouterHandlerFunc: cloud.CloudBatchJOBResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/deployment", RouterHandlerFunc: cloud.CloudDeploymentResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/daemonset", RouterHandlerFunc: cloud.CloudDaemonSetResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/replicaset", RouterHandlerFunc: cloud.CloudReplicasetResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/statefulset", RouterHandlerFunc: cloud.CloudStatefuleSetResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/pv", RouterHandlerFunc: cloud.CloudPVResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources/pvc", RouterHandlerFunc: cloud.CloudPVCResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/resources", RouterHandlerFunc: cloud.CloudResoucesHandler, RouterMethod: http.MethodGet},
		},
		InitFunc: []func() error{
			database.InitMySQLDB,
			k8s.InitKubernetesCluster,
		},
	}.WebServiceStart()

}
