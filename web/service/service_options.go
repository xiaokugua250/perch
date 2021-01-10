package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	database "perch/database/mysql"
	"perch/pkg/cluster/k8s"
	"perch/pkg/general/viperconf"

	log "github.com/sirupsen/logrus"
)

type OptionFunc func(options *WebServer)

func NewWebServerWithOptions(Name string, opts ...OptionFunc) WebServer {

	initFuncs := make(map[string]func(interface{}) error)

	webserver := WebServer{
		Name:   Name,
		Router: nil,
		//InitFuncConfigMaps: initFuncMaps,
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
				//DBConfig = "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
			}
		}

		//dsn := DBConfig

		database.MySQL_DB, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
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

		//dsn := DBConfig

		database.MySQL_DB, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
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

		database.MySQL_DB, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func WithKubernetesOptions(dbconfig interface{}) OptionFunc {
	return func(options *WebServer) {
		var (
			k8sClusterManager k8s.ClusterManager
			err               error
		)

		RunEnv := os.Getenv("RUN_ENV")
		if RunEnv == "" {
			RunEnv = "dev"
		}
		//	viper.AddConfigPath("E:\\WorksSpaces\\GoWorkSpaces\\perch\\configs\\"+RunEnv+"\\cluster_config")
		viper.AddConfigPath("E:\\WorksSpaces\\GoWorkSpaces\\perch\\configs\\dev\\cluster_config")
		viper.SetConfigName("kubernetes")
		//viper.SetConfigFile("kubernetes_cluster.config")
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalln(err)
		}
		err = viper.Unmarshal(&k8sClusterManager, func(config *mapstructure.DecoderConfig) {
			config.TagName = "yaml"
		})
		if err != nil {
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
		k8s.ClusterClientMap[k8sClusterManager.KubeConfig.ClusterName] = k8s.ClientSet{
			K8SClientSet:      clientset,
			K8sDynamitcClient: &dynamicClient,
		}
		k8s.K8SClientSet.K8SClientSet = clientset
		k8s.K8SClientSet.K8sDynamitcClient = &dynamicClient

	}
}
