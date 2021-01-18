package k8scloud

import (
	"context"
	"errors"
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	v1 "k8s.io/api/core/v1"

	batchV1 "k8s.io/api/batch/v1"
	v2alpha1 "k8s.io/api/batch/v2alpha1"
	betaV1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"net/http"
	"perch/pkg/cluster/k8scloud"
	"perch/web/metric"
	"perch/web/model"
)

/**
@cluster
@namespaces
*/
func CloudClustersHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			clusters []string
		)
		response.Kind = "cloud resources"

		for cluster, _ := range k8scloud.ClusterClientMap {
			clusters = append(clusters, cluster)
		}

		response.Code = http.StatusOK
		response.Spec = clusters
		response.Message = "k8s cluster namespaces"
		return nil
	})
}

/**
@cluster
@namespaces
*/
func CloudNameSpacesResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resultMap = make(map[string]interface{})
			cluster   string
			namespace string
		)
		response.Kind = "cloud resources"

		cluster = r.URL.Query().Get("cluster")
		namespace = r.URL.Query().Get("namespace")
		if cluster != "" {
			clientSet, ok := k8scloud.ClusterClientMap[cluster]
			if ok {

				ns, err := clientSet.K8sResourceListWithInformer(k8scloud.KubernetesNamespaces, namespace, labels.NewSelector())
				if err != nil {
					return err
				}
				namespaces, ok := ns.([]*v1.Namespace)
				if !ok {
					response.Code = http.StatusInternalServerError
					response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
					return err
				}
				resultMap[cluster] = namespaces
			} else {
				return errors.New(fmt.Sprintf("cluster %s not found....", cluster))
			}
		} else {
			for cluster, clientSet := range k8scloud.ClusterClientMap {
				ns, err := clientSet.K8sResourceListWithInformer(k8scloud.KubernetesNamespaces, namespace, labels.NewSelector())
				if err != nil {
					return err
				}
				namespaces, ok := ns.([]*v1.Namespace)
				if !ok {
					response.Code = http.StatusInternalServerError
					response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
					return err
				}
				resultMap[cluster] = namespaces

			}
		}

		response.Code = http.StatusOK
		response.Spec = resultMap
		response.Message = "k8s cluster namespaces"
		return nil
	})
}

/**
@cluster
@namespaces
*/
func CloudNodeResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resultMap = make(map[string]interface{})
			cluster   string
			namespace string
		)
		response.Kind = "cloud resources"

		cluster = r.URL.Query().Get("cluster")
		namespace = r.URL.Query().Get("namespace")
		fmt.Printf("%s,%s", cluster, namespace)
		if cluster != "" {
			clientSet, ok := k8scloud.ClusterClientMap[cluster]
			if ok {

				ns, err := clientSet.K8sResourceListWithInformer(k8scloud.KubernetesNode, namespace, labels.NewSelector())
				if err != nil {
					return err
				}
				nodes, ok := ns.([]*v1.Node)
				if !ok {
					response.Code = http.StatusInternalServerError
					response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
					return err
				}
				resultMap[cluster] = nodes
			} else {
				return errors.New(fmt.Sprintf("cluster %s not found....", cluster))
			}
		} else {
			for cluster, clientSet := range k8scloud.ClusterClientMap {
				ns, err := clientSet.K8sResourceListWithInformer(k8scloud.KubernetesNode, namespace, labels.NewSelector())
				if err != nil {
					return err
				}
				namespaces, ok := ns.([]*v1.Node)
				if !ok {
					response.Code = http.StatusInternalServerError
					response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
					return err
				}
				resultMap[cluster] = namespaces

			}
		}

		response.Code = http.StatusOK
		response.Spec = resultMap
		response.Message = "k8s cluster namespaces"
		return nil
	})
}

/**
@cluster
//todo 工作负载显示集群信息
*/
func CloudWorkloadHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			cluster string

			workload struct {
				Master   []string `json:"master"`
				MatserMachine interface{} `json:"matser_machine"`
				Worker   []string `json:"worker"`
				WorkerMachine interface{}`json:"worker_machine"`
				Resource struct {
					Avaliable struct {
						CPU         int64 `json:"cpu"`
						Storage     int64 `json:"storage"`
						StorageEphemeral     int64 `json:"storage_ephemeral"`
						Memory      int64 `json:"memory"`
						NvidiaGPU   int   `json:"nvidia_gpu"`
					} `json:"avaliable"`
					Used struct {
						CPU         int64 `json:"cpu"`
						Storage     int64 `json:"storage"`
						Memory      int64 `json:"memory"`
						StorageEphemeral     int64 `json:"storage_ephemeral"`

						NvidiaGPU   int   `json:"nvidia_gpu"`
					} `json:"used"`
				} `json:"resource"`
			}
		)

		response.Kind = "cloud resources"
		cluster = mux.Vars(r)["cluster"]

		if cluster != "" {
			clientSet, ok := k8scloud.ClusterClientMap[cluster]
			if ok {
				ns, err := clientSet.K8sResourceListWithInformer(k8scloud.KubernetesNode, "", labels.NewSelector())
				if err != nil {
					return err
				}
				nodes, ok := ns.([]*v1.Node)
				if !ok {
					response.Code = http.StatusInternalServerError
					response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
					return err
				}
				for _,node:= range nodes{
					if _,ok := node.Labels["node-role.kubernetes.io/master"];ok{
						workload.Master=append(workload.Master,node.Name)
						mastermachine := make(map[string]string)
						mastermachine["arch"]=node.Labels["beta.kubernetes.io/arch"]
						mastermachine["os"]=node.Labels["beta.kubernetes.io/os"]
						workload.MatserMachine=node.Status.NodeInfo

					}else { //todo 对于混合集群，这种处理方式错误
						workload.Worker=append(workload.Worker,node.Name)
						workermachine := make(map[string]string)
						workermachine["arch"]=node.Labels["beta.kubernetes.io/arch"]
						workermachine["os"]=node.Labels["beta.kubernetes.io/os"]
						workload.WorkerMachine=workermachine
					}
					workload.Resource.Avaliable.CPU+=node.Status.Allocatable.Cpu().Value()
					workload.Resource.Avaliable.Memory+=node.Status.Allocatable.Memory().Value()/1024/1024/1024
					workload.Resource.Avaliable.Storage+=node.Status.Allocatable.Storage().Value()
					workload.Resource.Avaliable.StorageEphemeral+=node.Status.Allocatable.StorageEphemeral().Value()/1024/1024/1024

				}
				//resultMap[cluster]=nodes

			} else {
				return errors.New(fmt.Sprintf("cluster %s not found....", cluster))
			}
		}

		response.Code = http.StatusOK
		response.Spec = workload
		response.Message = "k8s cluster namespaces"
		return nil
	})
}

/**
@cluster
@namespaces
@type
	@type ref to pod,job,deployment,service etc...
*/
func CloudWorkloadResourcesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resultMap        = make(map[string]interface{})
			cluster          string
			namespace        string
			workloadResource string
			result           interface{}
		)
		response.Kind = "cloud resources"

		cluster = mux.Vars(r)["clutser"]
		namespace = mux.Vars(r)["namespace"]
		workloadResource = mux.Vars(r)["type"]
		if cluster != "" {
			clientSet, ok := k8scloud.ClusterClientMap[cluster]
			if ok {
				result_obj, err := clientSet.K8sResourceListWithInformer(workloadResource, namespace, labels.NewSelector())
				if err != nil {
					return err
				}
				switch workloadResource {
				case k8scloud.KubernetesDeployment:
					result, ok = result_obj.([]*betaV1.Deployment)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesJOB:
					result, ok = result_obj.([]*batchV1.Job)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesCronjob:
					result, ok = result_obj.([]*v2alpha1.CronJob)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesDaemonset:
					result, ok = result_obj.([]*betaV1.DaemonSet)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesPod:
					result, ok = result_obj.([]*v1.Pod)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesPvc:
					result, ok = result_obj.([]*v1.PersistentVolumeClaim)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesPv:
					result, ok = result_obj.([]*v1.PersistentVolume)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				case k8scloud.KubernetesService:
					result, ok = result_obj.([]*v1.Service)
					if !ok {
						response.Code = http.StatusInternalServerError
						response.Message = fmt.Sprintf("marsh result to v1.namespaces failed...")
						return err
					}
				}

				resultMap[cluster] = result
			} else {
				return errors.New(fmt.Sprintf("cluster %s not found....", cluster))
			}
		}

		response.Code = http.StatusOK
		response.Spec = resultMap
		response.Message = "k8s cluster namespaces"
		return nil
	})
}

/*
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
*/
