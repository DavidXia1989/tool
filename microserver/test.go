package template

var Test = `package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/client/grpc"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/registry/etcd"
    opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
    "github.com/opentracing/opentracing-go"
    "{{.Name}}/common"
    "{{.Name}}/conf"
    "{{.Name}}/proto/example"
    "time"
)

func init(){
    conf.LoadAppCenterConf()
}

func main() {
    // 初始化链路追踪
    t, io, err := common.NewTracer(conf.Conf.Micro.TracerClientName, conf.Conf.Micro.TracerAddr)
    if err != nil {
        // 链路追踪初始化失败
    }
    defer io.Close()
    opentracing.SetGlobalTracer(t)

    // rpc链接配置（客户端版）
    client.NewClient = grpc.NewClient
    c := client.NewClient(
        client.PoolSize(1),
        client.Retries(1),
        client.DialTimeout(time.Second*2),
        client.Registry(etcd.NewRegistry(registry.Addrs(conf.Conf.Micro.Registry))),
        client.Wrap(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
    )

    // 绑定业务客户端 至 rpc
    exampleClient := example.NewExampleService(conf.Conf.Micro.ProjectName, c)

    // 发送请求
    rep, err := exampleClient.Hello(context.TODO(), &example.HelloReq{Name: "World"})
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(rep.Msg)

}
`
