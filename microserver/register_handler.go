package template

var RegisterHandler = `
package handler

import (
"code.zm.shzhanmeng.com/go-common/logging"
"fmt"
"github.com/micro/go-micro/v2"
"github.com/micro/go-micro/v2/registry"
"github.com/micro/go-micro/v2/registry/etcd"
ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
"github.com/opentracing/opentracing-go"
"go.uber.org/zap"
"{{.Name}}/common"
"{{.Name}}/conf"
service2 "microServerExample/domain/service"
"{{.Name}}/proto/example"
)

var MicroServer micro.Service

func RegistryHandler(){
	//创建链路追踪实例

	t,io,err := common.NewTracer(conf.Conf.Micro.TracerServerName, conf.Conf.Micro.TracerAddr)
	if err != nil {
		logging.ZapLogger.Error(err.Error())
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	grpcAddr := conf.Conf.Micro.GrpcAddr + ":" + conf.Conf.Micro.GrpcPort

	MicroServer = micro.NewService(
		micro.Name(conf.Conf.Micro.ProjectName),
		micro.Address(grpcAddr),
		micro.Registry(etcd.NewRegistry(registry.Addrs(conf.Conf.Micro.Registry))),
		// 限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(conf.Conf.Micro.Qps)),
		// 链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	MicroServer.Init()

	var errs  []error
	//注册服务
	errs = append(errs, example.RegisterExampleHandler(MicroServer.Server(),&Example{
		ExampleService: service2.NewExampleApi(),
	}))

	for k:= range errs{
		if errs[k] !=nil {
			logging.ZapLogger.Info("业务注册失败",zap.Error(errs[k]))
			fmt.Println(errs[k].Error())
		}
	}

	if err = MicroServer.Run(); err != nil {
		logging.ZapLogger.Info("micro grpc 启动失败",zap.Error(err))
		fmt.Println(err)
	}
}`

