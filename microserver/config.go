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

// micro配置中心
type Micro struct {
	TracerServerName string	` + "`" + `yaml:"tracer_server_name"` + "`" + `
	TracerClientName string	` + "`" + `yaml:"tracer_client_name"` + "`" + `
	TracerAddr	string	` + "`" + `yaml:"tracer_addr"` + "`" + `
	GrpcAddr	string	` + "`" + `yaml:"grpc_addr"` + "`" + `
	GrpcPort	string	` + "`" + `yaml:"grpc_port"` + "`" + `
	ProjectName	string	` + "`" + `yaml:"project_name"` + "`" + `
	Registry	string	` + "`" + `yaml:"registry"` + "`" + `
	Qps	string	` + "`" + `yaml:"qps"` + "`" + `
}

type conf struct {
	Micro			Micro       ` + "`" + `yaml:"micro"` + "`" + `
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
