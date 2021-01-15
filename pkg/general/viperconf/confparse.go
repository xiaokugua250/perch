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

var (
	DefaultDevConfigDir = "configs/dev/web_config/" //DEV_CONFIG_DIR
	DefaultProConfigDir = "configs/pro/web_config/" //PRO_CONFIG_DIR
)

func init() {
	if runtime.GOOS != "windows" {
		DefaultDevConfigDir = "/configs/dev/web_config/" //DEV_CONFIG_DIR
		DefaultProConfigDir = "/configs/pro/web_config/" //PRO_CONFIG_DIR
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
type GenWebConfig struct {
	WebConfig WebServerConfig `yaml:"web-server"`
}
type DataBaseConfig struct {
}

var WebServiceConfig *GenWebConfig

/**
解析yaml形式config文件
*/
func InitYamlconfig(configfile string, result *interface{}) error {
	_, err := os.Stat(configfile)
	if os.IsNotExist(err) {
		return err
	}
	conifgyaml, err := ioutil.ReadFile(configfile)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(conifgyaml, &result); err != nil {
		return err
	}
	return nil
}

func InitGenWebConfig(configfile string) error {
	var (
		err error
	)
	if os.Getenv("DEV_CONFIG_DIR") != "" {
		DefaultDevConfigDir = os.Getenv("DEV_CONFIG_DIR")
	}
	if os.Getenv("PRO_CONFIG_DIR") != "" {
		DefaultProConfigDir = os.Getenv("DEV_CONFIG_DIR")
	}
	if os.Getenv("PRO_ENV") != "" {
		configfile = DefaultProConfigDir + configfile + ".yaml"
	} else {
		configfile = DefaultDevConfigDir + configfile + ".yaml"

	}
	_, err = os.Stat(configfile)
	if os.IsNotExist(err) {
		return err
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
