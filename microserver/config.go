package template

var Conf = `package conf

import (
	"{{.Name}}/kernel"
)

// 读取配置中心
type AppCenter struct {
	AppHost string	` + "`" + `yaml:"app_host"` + "`" + `
	AppUrl	string	` + "`" + `yaml:"app_url"` + "`" + `
}
type conf struct {
	AppCenter		AppCenter	` + "`" + `yaml:"app_center"` + "`" + `
}

var (
	Conf = &conf{}
)

func LoadAppCenterConf(){
	err := kernel.GetConfig(Conf)
	if err != nil {
		panic("加载设置动态配置错误")
	}

}`
