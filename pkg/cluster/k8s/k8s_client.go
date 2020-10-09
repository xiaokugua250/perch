package k8s

import (
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "perch/pkg/log"
)

type K8SClusterManager struct {
	KubeClusterName string `yaml:"cluster_name"`
	KubeConfig      string `yaml:"kubeconfig"`
}

type ClientSet struct {
	K8SClientSet *kubernetes.Clientset
}

var (
	ClusterClientMap = make(map[string]*kubernetes.Clientset)
	K8SClientSet     = ClientSet{}
)

/**
初始化k8s集群
*/
func (k8sClusterManager *K8SClusterManager) InitK8SClusterClient() error {
	config, err := clientcmd.BuildConfigFromFlags("", k8sClusterManager.KubeConfig)
	if err != nil {
		log.Fatalln(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	ClusterClientMap[k8sClusterManager.KubeClusterName] = clientset
	return nil

}
