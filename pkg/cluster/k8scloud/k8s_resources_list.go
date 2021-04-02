package k8scloud

import (
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	appv1 "k8s.io/client-go/informers/apps/v1"
	batchv2 "k8s.io/client-go/informers/batch/v2alpha1"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	_ "perch/pkg/log"
	"time"
)

/**
采用list and watch机制获取资源列表
*/
func (k8sClientSet *ClientSet) K8sResourceListWithInformer(resouceType string, namespaces string, selector labels.Selector) (interface{}, error) {
	var (
		informer interface{}
	)
	stopChan := make(chan struct{})
	factory := informers.NewSharedInformerFactoryWithOptions(k8sClientSet.K8SClientSet, 15*time.Second, informers.WithNamespace(namespaces))
	switch resouceType {
	case KubernetesNode:
		informer = factory.Core().V1().Nodes()
		//nodeInformer:= factory.Core().V1().Nodes()
		go informer.(v1.NodeInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.NodeInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.NodeInformer).Lister().List(selector)
	case KubernetesNamespaces: //todo 只获取特定namespaces
		informer = factory.Core().V1().Namespaces()
		go informer.(v1.NamespaceInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.NamespaceInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.NamespaceInformer).Lister().List(selector)
	case K8sResourceConfigmap:
		informer = factory.Core().V1().ConfigMaps()
		go informer.(v1.ConfigMapInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.ConfigMapInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.ConfigMapInformer).Lister().List(selector)
	case KubernetesService:
		informer = factory.Core().V1().Services()
		go informer.(v1.ServiceInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.ServiceInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.ServiceInformer).Lister().List(selector)
	case K8sResourceServiceaccount:
		informer = factory.Core().V1().ServiceAccounts()
		go informer.(v1.ServiceAccountInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.ServiceAccountInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.ServiceAccountInformer).Lister().List(selector)
	case KubernetesPod:
		informer = factory.Core().V1().Pods()
		go informer.(v1.PodInformer).Informer().Run(stopChan)

		if !cache.WaitForCacheSync(stopChan, informer.(v1.PodInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.PodInformer).Lister().List(selector)
	case KubernetesJOB: //todo 需要重新处理
		Jobinformer := factory.Batch().V1().Jobs().Informer()
		go Jobinformer.Run(stopChan)
		//go factory.Batch().V1().Jobs().Informer().Run(stopChan)
		//go informer.(batchv1.JobInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, Jobinformer.HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}

		return factory.Batch().V1().Jobs().Lister().List(selector)
	//	return informer.(batchv1.JobInformer).Lister().List(selector)
	case KubernetesCronjob: //todo 需要重新处理
		BatchJobinformer := factory.Batch().V2alpha1().CronJobs().Informer()
		//	go informer.(batchv2.CronJobInformer).Informer().Run(stopChan)
		go BatchJobinformer.Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, BatchJobinformer.HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}

		return informer.(batchv2.CronJobInformer).Lister().List(selector)
	case KubernetesDeployment:
		informer = factory.Apps().V1().Deployments()
		go informer.(appv1.DeploymentInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(appv1.DeploymentInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(appv1.DeploymentInformer).Lister().List(selector)

	case KubernetesDaemonset:
		informer = factory.Apps().V1().DaemonSets()
		go informer.(appv1.DaemonSetInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(appv1.DaemonSetInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(appv1.DaemonSetInformer).Lister().List(selector)

	case K8sResourceReplicaset:
		informer = factory.Apps().V1().ReplicaSets()
		go informer.(appv1.ReplicaSetInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(appv1.ReplicaSetInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(appv1.ReplicaSetInformer).Lister().List(selector)
	case K8sResourceStatefulset:
		informer = factory.Apps().V1().StatefulSets()
		go informer.(appv1.StatefulSetInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(appv1.StatefulSetInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(appv1.StatefulSetInformer).Lister().List(selector)
	case KubernetesPv:
		informer = factory.Core().V1().PersistentVolumes()
		go informer.(v1.PersistentVolumeInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.PersistentVolumeInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("Timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.PersistentVolumeInformer).Lister().List(selector)
	case KubernetesPvc:
		informer = factory.Core().V1().PersistentVolumeClaims()
		go informer.(v1.PersistentVolumeClaimInformer).Informer().Run(stopChan)
		//go informer.Informer().Run(stopChan)
		if !cache.WaitForCacheSync(stopChan, informer.(v1.PersistentVolumeClaimInformer).Informer().HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return nil, errors.New("timed out waiting for caches to sync,failed to get  resources information...")
		}
		return informer.(v1.PersistentVolumeClaimInformer).Lister().List(selector)
	default:
		return nil, errors.New("resouce type not support now...")
	}

}
