package template

var MainFunc = `package main

import (
	"{{.Name}}/kernel"
	"{{.Name}}/conf"
)

func init(){
	// 动态加载配置，加载配置中心结构
	conf.LoadAppCenterConf()
}

func main() {
	//ctx := kernel.SetSignal()
	kernel.SetupRedis()
	kernel.SetupMysql()
	//go kernel.SetupCron(ctx)
	kernel.SetupGrpc()

}`
