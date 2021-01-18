package k8scloud

import (
	"context"
	"fmt"
	_ "fmt"
	app "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"net/http"
	"perch/pkg/cluster/k8scloud"
	"perch/web/metric"
	"perch/web/model"
)

func CloudNameSpacesResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		fmt.Printf("==>%v")
		fmt.Println(k8scloud.K8SClientSet)
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceNamespaces, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		namespaces, ok := result.([]*v1.Namespace)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = namespaces
		response.Message = "k8s cluster namespaces"
		return nil
	})
}

func CloudNodeResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceNode, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		nodes, ok := result.([]*v1.Node)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.nodes failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = nodes
		response.Message = "k8s cluster namespaces"
		return nil
	})
}
func CloudConfigMapResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceConfigmap, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		configmaps, ok := result.([]*v1.ConfigMap)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.configmaps failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = configmaps
		response.Message = "k8s cluster configmaps"
		return nil
	})
}
func CloudServiceAccountResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceServiceaccount, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		serviceaccount, ok := result.([]*v1.ServiceAccount)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.serviceaccount failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = serviceaccount
		response.Message = "k8s cluster serviceaccount"
		return nil
	})
}
func CloudPODResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourcePod, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		pods, ok := result.([]*v1.Pod)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.pod failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = pods
		response.Message = "k8s cluster pods"
		return nil
	})
}
func CloudJOBResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8S_RESOURCE_JOB, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		jobs, ok := result.([]*batchv1.Job)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.job failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = jobs
		response.Message = "k8s cluster jobs"
		return nil
	})
}
func CloudBatchJOBResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceBatchjob, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		crobJobs, ok := result.([]*batchv1beta1.CronJob)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.crobjb failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = crobJobs
		response.Message = "k8s cluster crobJobs"
		return nil
	})
}
func CloudServiceResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceService, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		services, ok := result.([]*v1.Service)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.service failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = services
		response.Message = "k8s cluster services"
		return nil
	})
}
func CloudDeploymentResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceDeployment, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		deployments, ok := result.([]*app.Deployment)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.deployments failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = deployments
		response.Message = "k8s cluster deployments"
		return nil
	})
}
func CloudDaemonSetResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceDaemonset, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		daemonsets, ok := result.([]*app.DaemonSet)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.deamonset failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = daemonsets
		response.Message = "k8s cluster daemonsets"
		return nil
	})
}
func CloudReplicasetResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceReplicaset, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		replicasets, ok := result.([]*app.ReplicaSet)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.replicaset failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = replicasets
		response.Message = "k8s cluster replicasets"
		return nil
	})
}
func CloudStatefuleSetResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourceStatefulset, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		statefulsets, ok := result.([]*app.StatefulSet)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to app.statefuleset failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = statefulsets
		response.Message = "k8s cluster statefulsets"
		return nil
	})
}
func CloudPVResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourcePv, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		pv, ok := result.([]*v1.PersistentVolume)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.pv failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = pv
		response.Message = "k8s cluster pv"
		return nil
	})
}
func CloudPVCResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
		)
		response.Kind = "cloud resources"
		result, err := k8scloud.K8SClientSet.K8sResourceListWithInformer(k8scloud.K8sResourcePvc, "", labels.NewSelector())
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		pvc, ok := result.([]*v1.PersistentVolumeClaim)
		if !ok {
			response.Code = http.StatusInternalServerError
			response.Message = fmt.Sprintf("marsh result to v1.pvc failed...")
			return err
		}
		response.Code = http.StatusOK
		response.Spec = pvc
		response.Message = "k8s cluster pvc"
		return nil
	})
}
