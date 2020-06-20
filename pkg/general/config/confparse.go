package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

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
