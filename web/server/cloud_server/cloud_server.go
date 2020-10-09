package main

import (
	"net/http"
	cloud "perch/api/cloud_api"
	database "perch/database/mysql"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "plat-cloud",

		Router: []service.WebRouter{
			{RouterPath: "/cloud/resources/namespaces", RouterHandlerFunc: cloud.CloudNameSpacesResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/nodes", RouterHandlerFunc: cloud.CloudNodeResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/service", RouterHandlerFunc: cloud.CloudServiceResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/configmap", RouterHandlerFunc: cloud.CloudConfigMapResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/serviceaccount", RouterHandlerFunc: cloud.CloudServiceAccountResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/pod", RouterHandlerFunc: cloud.CloudPODResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/job", RouterHandlerFunc: cloud.CloudJOBResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/batchjob", RouterHandlerFunc: cloud.CloudBatchJOBResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/deployment", RouterHandlerFunc: cloud.CloudDeploymentResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/daemonset", RouterHandlerFunc: cloud.CloudDaemonSetResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/replicaset", RouterHandlerFunc: cloud.CloudReplicasetResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/statefulset", RouterHandlerFunc: cloud.CloudStatefuleSetResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/pv", RouterHandlerFunc: cloud.CloudPVResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources/pvc", RouterHandlerFunc: cloud.CloudPVCResoucesHandler, RouterMethod: http.MethodGet},
			{RouterPath: "/cloud/resources", RouterHandlerFunc: cloud.CloudResoucesHandler, RouterMethod: http.MethodGet},
		},
		InitFunc: []func() error{
			database.InitMySQLDB,
		},
	}.WebServiceStart()

}
