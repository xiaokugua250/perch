package k8s

import (
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	_ "perch/pkg/log"
)


type ClusterManager struct {
	KubeConfig      KUBEConfig`yaml:"kubeconfig"`
}
type KUBEConfig struct {
	ClusterName string `yaml:"cluster_name"`
	ConfigFile string `yaml:"config_file"`
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
func (k8sClusterManager *ClusterManager) InitK8SClusterClient() error {

	config, err := clientcmd.BuildConfigFromFlags("", k8sClusterManager.KubeConfig.ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	ClusterClientMap[k8sClusterManager.KubeConfig.ClusterName] = clientset
	return nil

}


/**
初始化k8s集群
*/
func  InitKubernetesCluster() error {
	var (
		k8sClusterManager  ClusterManager
		err error
	)

	RunEnv := os.Getenv("RUN_ENV")
	if RunEnv == ""{
		RunEnv="dev"
	}
//	viper.AddConfigPath("E:\\WorksSpaces\\GoWorkSpaces\\perch\\configs\\"+RunEnv+"\\cluster_config")
	viper.AddConfigPath("E:\\WorksSpaces\\GoWorkSpaces\\perch\\configs\\dev\\cluster_config")
	viper.SetConfigName("kubernetes")
	//viper.SetConfigFile("kubernetes_cluster.config")
	err = viper.ReadInConfig()
	if err!= nil{
		return err
	}
	err =viper.Unmarshal(&k8sClusterManager, func(config *mapstructure.DecoderConfig) {
		config.TagName="yaml"
	})
	if err!= nil{
		return err
	}
	config, err := clientcmd.BuildConfigFromFlags("", k8sClusterManager.KubeConfig.ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	ClusterClientMap[k8sClusterManager.KubeConfig.ClusterName] = clientset
	K8SClientSet.K8SClientSet=clientset
	return nil

}
