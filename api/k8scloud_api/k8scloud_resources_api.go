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
				Master        []string    `json:"master"`
				MatserMachine interface{} `json:"matser_machine"`
				Worker        []string    `json:"worker"`
				WorkerMachine interface{} `json:"worker_machine"`
				Resource      struct {
					Avaliable struct {
						CPU              int64 `json:"cpu"`
						Storage          int64 `json:"storage"`
						StorageEphemeral int64 `json:"storage_ephemeral"`
						Memory           int64 `json:"memory"`
						NvidiaGPU        int   `json:"nvidia_gpu"`
					} `json:"avaliable"`
					Used struct {
						CPU              int64 `json:"cpu"`
						Storage          int64 `json:"storage"`
						Memory           int64 `json:"memory"`
						StorageEphemeral int64 `json:"storage_ephemeral"`

						NvidiaGPU int `json:"nvidia_gpu"`
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
				for _, node := range nodes {
					if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
						workload.Master = append(workload.Master, node.Name)
						mastermachine := make(map[string]string)
						mastermachine["arch"] = node.Labels["beta.kubernetes.io/arch"]
						mastermachine["os"] = node.Labels["beta.kubernetes.io/os"]
						workload.MatserMachine = node.Status.NodeInfo

					} else { //todo 对于混合集群，这种处理方式错误
						workload.Worker = append(workload.Worker, node.Name)
						workermachine := make(map[string]string)
						workermachine["arch"] = node.Labels["beta.kubernetes.io/arch"]
						workermachine["os"] = node.Labels["beta.kubernetes.io/os"]
						workload.WorkerMachine = workermachine
					}
					workload.Resource.Avaliable.CPU += node.Status.Allocatable.Cpu().Value()
					workload.Resource.Avaliable.Memory += node.Status.Allocatable.Memory().Value() / 1024 / 1024 / 1024
					workload.Resource.Avaliable.Storage += node.Status.Allocatable.Storage().Value()
					workload.Resource.Avaliable.StorageEphemeral += node.Status.Allocatable.StorageEphemeral().Value() / 1024 / 1024 / 1024

				}

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
