package service

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	database "perch/database/mysql"
	"perch/pkg/cluster/k8scloud"
	"perch/pkg/general/viperconf"

	log "github.com/sirupsen/logrus"
)

type OptionFunc func(options *WebServer)

func NewWebServerWithOptions(Name string, opts ...OptionFunc) WebServer {

	initFuncs := make(map[string]func(interface{}) error)
	webserver := WebServer{
		Name:   Name,
		Router: nil,

		InitFuncs: initFuncs,
		CleanFunc: nil,
	}
	webserver.Init()
	for _, o := range opts {
		o(&webserver)
	}

	return webserver
}

func WithMySQLDBOptions(dbconfig interface{}) OptionFunc {
	return func(options *WebServer) {
		var (
			err      error
			DBConfig string
		)
		if dbconfig, ok := dbconfig.(string); ok {

			if dbconfig != "" { //"genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local
				DBConfig = dbconfig
			} else {

				DBConfig = viperconf.WebServiceConfig.WebConfig.ServerDB.DBConnURL
			}
		}
		database.MysqlDb, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

	}
}
func WithRedisOptions(redisConfig interface{}) OptionFunc {
	return func(options *WebServer) {
		var (
			err      error
			DBConfig string
		)
		if dbconfig, ok := redisConfig.(string); ok {
			if dbconfig != "" {
				DBConfig = dbconfig
			} else {
				DBConfig = "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
			}
		}

		database.MysqlDb, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func WithETCDOptions(etcdConfig interface{}) OptionFunc {
	return func(options *WebServer) {
		var (
			err      error
			DBConfig string
		)
		if dbconfig, ok := etcdConfig.(string); ok {
			if dbconfig != "" {
				DBConfig = dbconfig
			} else {
				DBConfig = "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
			}
		}

		//dsn := DBConfig

		database.MysqlDb, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func WithClustersOptions(cluster interface{}) OptionFunc {
	return func(options *WebServer) {
		var (
			k8sClusterManager k8scloud.ClusterManager
			err               error
			clusterConfig     string
		)
		if clusterConfigStr, ok := cluster.(string); ok {
			if clusterConfig != "" { //"genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local
				clusterConfig = clusterConfigStr
			} else {
				if os.Getenv("RUN_ENV") != "" {
					clusterConfig = viperconf.DefaultconfigsDir + os.Getenv("RUN_ENV") + "/cluster_config/cluster.yaml"
				} else {
					clusterConfig = viperconf.DefaultconfigsDir + "/dev/cluster_config/cluster.yaml"
				}

			}
		}

		conifgyaml, err := ioutil.ReadFile(clusterConfig)
		if err != nil {
			log.Fatalln(err)
		}
		if err = yaml.Unmarshal(conifgyaml, &viperconf.ClustersConfigurations); err != nil {
			log.Fatalln(err)
		}

		config, err := clientcmd.BuildConfigFromFlags("", k8sClusterManager.KubeConfig.ConfigFile)
		if err != nil {
			log.Fatalln(err)
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalln(err)
		}
		dynamicClient, err := dynamic.NewForConfig(config)
		if err != nil {
			log.Fatalln(err)
		}
		k8scloud.ClusterClientMap[k8sClusterManager.KubeConfig.ClusterName] = k8scloud.ClientSet{
			K8SClientSet:      clientset,
			K8sDynamitcClient: &dynamicClient,
		}


	}
}

func WithKubernetesOptions(configfile interface{}) OptionFunc {
	return func(options *WebServer) {
		var (
			err           error
			clusterConfig string
		)
		if clusterConfigStr, ok := configfile.(string); ok {
			if clusterConfig != "" { //"genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local
				clusterConfig = clusterConfigStr
			} else {
				if os.Getenv("RUN_ENV") != "" {
					clusterConfig = viperconf.DefaultconfigsDir + os.Getenv("RUN_ENV") + "/cluster_config/cluster.yaml"
				} else {
					clusterConfig = viperconf.DefaultconfigsDir + "/dev/cluster_config/cluster.yaml"
				}
			}
		}

		conifgyaml, err := ioutil.ReadFile(clusterConfig)
		if err != nil {
			log.Fatalln(err)
		}
		if err = yaml.Unmarshal(conifgyaml, &viperconf.ClustersConfigurations); err != nil {
			log.Fatalln(err)
		}

		for _, cluster := range viperconf.ClustersConfigurations.Clusters {
			if cluster.ClusterConfig.ClusterType == viperconf.CLUSTER_TYPE_KUBERNETES {
				fmt.Printf("%+s", cluster.ClusterConfig.ClusterFile)
				config, err := clientcmd.BuildConfigFromFlags("", cluster.ClusterConfig.ClusterFile)
				if err != nil {
					log.Fatalln(err)
				}
				clientset, err := kubernetes.NewForConfig(config)
				if err != nil {
					log.Fatalln(err)
				}
				dynamicClient, err := dynamic.NewForConfig(config)
				if err != nil {
					log.Fatalln(err)
				}
				k8scloud.ClusterClientMap[cluster.ClusterConfig.ClusterName] = k8scloud.ClientSet{
					K8SClientSet:      clientset,
					K8sDynamitcClient: &dynamicClient,
				}
			}
		}

	}

}
