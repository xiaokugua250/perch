package k8scloud

import (
	_ "perch/pkg/log"
)

const (
	KubernetesNode            = "NODE"
	K8sResourceConfigmap      = "CONFIGMAP"
	KubernetesNamespaces      = "NAMESPACES"
	K8sResourceServiceaccount = "SERVICEACCOUNT"
	KubernetesPod             = "POD"
	KubernetesJOB             = "JOB"
	KubernetesCronjob         = "CronJob"
	KubernetesService         = "SERVICE"
	KubernetesDeployment      = "DEPLOYMENT"
	KubernetesDaemonset       = "DAEMONSET"
	K8sResourceReplicaset     = "REPLICASET"
	K8sResourceStatefulset    = "STATEFULESET"
	KubernetesPv              = "PersistentVolume"
	KubernetesPvc            = "PersistentVolumeClaim"
)
