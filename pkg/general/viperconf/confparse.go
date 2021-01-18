package viperconf

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
)

/*
通用web服务配置
`
web-server:
  ip: 127.0.0.1
  port: 8080
  database:
    dbconn_url: "root:mysqladmin@tcp(127.0.0.1:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
    cacheconn_url: "127.0.0.1:6379"
`
*/

/*集群配置文件

clusters_config:
  - cluster:
      name:
	  type:
      config_file:
      config:

*/
var (
	DefaultconfigsDir               string //默认配置文件目录
	DefaultClusterconfigurationfile string
)

func init() {

	if os.Getenv("CONFIG_DIR") != "" {
		DefaultconfigsDir = os.Getenv("DEV_CONFIG_DIR")
	} else {
		DefaultconfigsDir = "configs"
		if runtime.GOOS != "windows" {
			DefaultconfigsDir = "/configs"
		}
	}

}

type WebServerConfig struct {
	ServerIP   string `yaml:"ip"`
	ServerPort int    `yaml:"port"`
	ServerDB   struct {
		DBConnURL string `yaml:"dbconn_url"`
		CacheURL  string `yaml:"cacheconn_url"`
	} `yaml:"database"`
}
type GeneralWebConfig struct {
	WebConfig WebServerConfig `yaml:"web-server"`
}
type DataBaseConfig struct {
}

var WebServiceConfig *GeneralWebConfig

func InitGeneralWebConfig(configfile string) error {
	var (
		err error
	)

	if os.Getenv("RUN_ENV") != "" {
		configfile = DefaultconfigsDir + os.Getenv("RUN_ENV") + "/web_config/" + configfile + ".yaml"
	} else {
		configfile = DefaultconfigsDir + "/dev/web_config/" + configfile + ".yaml"
	}
	conifgyaml, err := ioutil.ReadFile(configfile)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(conifgyaml, &WebServiceConfig); err != nil {
		return err
	}
	return nil
}

//***********clusters config *********************

type Cluster struct {
	ClusterConfig struct {
		ClusterName string `yaml:"cluster_name"`
		ClusterType string `yaml:"cluster_type"`
		ClusterFile string `yaml:"config_file"`
		//ClusterConfig string `yaml:"config"`
	} `yaml:"cluster"`
}

type ClusterConfiguration struct {
	Clusters []Cluster `yaml:"clusters_configiruation"`
}

var ClustersConfigurations *ClusterConfiguration

const (
	CLUSTER_TYPE_KUBERNETES = "kubernetes"
	CLUSTER_TYPE_SLURM      = "slurm"
)

type App struct {
	Address string
	Static  string
	Log     string
}

type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

type Configuration struct {
	App App
	Db  Database
}

var config *Configuration
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}
