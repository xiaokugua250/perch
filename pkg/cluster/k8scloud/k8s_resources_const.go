package k8scloud

import (
	_ "perch/pkg/log"
)

const (
	K8sResourceNode           = "NODE"
	K8sResourceConfigmap      = "CONFIGMAP"
	K8sResourceNamespaces     = "NAMESPACES"
	K8sResourceServiceaccount = "SERVICEACCOUNT"
	K8sResourcePod            = "POD"
	K8S_RESOURCE_JOB          = "JOB"
	K8sResourceBatchjob       = "BATCHJOB"
	K8sResourceService        = "SERVICE"
	K8sResourceDeployment     = "DEPLOYMENT"
	K8sResourceDaemonset      = "DAEMONSET"
	K8sResourceReplicaset     = "REPLICASET"
	K8sResourceStatefulset    = "STATEFULESET"
	K8sResourcePv             = "PersistentVolume"
	K8sResourcePvc            = "PersistentVolumeClaim"
)
