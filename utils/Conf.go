package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ConfUtil ...
type ConfUtil struct{}

// BindYamlConf 绑定配置
func BindYamlConf(configObject interface{}, confFile string) error {
	// 加载db配置文件及初始化
	dbYaml, err := ioutil.ReadFile(confFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(dbYaml, configObject)
	if err != nil {
		return err
	}
	return nil
}
