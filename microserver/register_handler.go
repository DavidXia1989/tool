package template

var RegisterHandler = `package handler

import (
	"code.zm.shzhanmeng.com/go-common/logging"
	"fmt"
	"go.uber.org/zap"
	service2 "{{.Name}}/domain/service"
	"{{.Name}}/kernel"
	"{{.Name}}/proto/example"
)

func RegistryHandler(){
	var errs  []error
	fmt.Println(kernel.MicroServer)
	//注册服务
	errs = append(errs, example.RegisterExampleHandler((*kernel.MicroServer).Server(),&Example{
		ExampleService: service2.NewExampleApi(),
	}))

	for k:= range errs{
		if errs[k] !=nil {
			logging.ZapLogger.Info("业务注册失败",zap.Error(errs[k]))
			fmt.Println(errs[k].Error())
		}
	}
}`
