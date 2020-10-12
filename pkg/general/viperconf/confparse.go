package viperconf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
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
	DEFAULT_DEV_CONFIG_DIR = "configs/dev/web_config/" //DEV_CONFIG_DIR
	DEFAULT_PRO_CONFIG_DIR = "configs/pro/web_config/" //PRO_CONFIG_DIR
)

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
		DEFAULT_DEV_CONFIG_DIR = os.Getenv("DEV_CONFIG_DIR")
	}
	if os.Getenv("PRO_CONFIG_DIR") != "" {
		DEFAULT_PRO_CONFIG_DIR = os.Getenv("DEV_CONFIG_DIR")
	}
	if os.Getenv("PRO_ENV") != "" {
		configfile = DEFAULT_PRO_CONFIG_DIR + configfile + ".yaml"
	} else {
		configfile = DEFAULT_DEV_CONFIG_DIR + configfile + ".yaml"

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
