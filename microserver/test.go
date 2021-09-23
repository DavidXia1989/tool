package template

var Test = `
package main

import (
    "context"
    "fmt"
    "{{.Name}}/conf"
    "{{.Name}}/proto/example"
    "time"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/client/grpc"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
    client.NewClient = grpc.NewClient
    c := client.NewClient(
        client.PoolSize(1),
        client.Retries(1),
        client.DialTimeout(time.Second*2),
        client.Registry(etcd.NewRegistry(registry.Addrs(conf.Conf.Micro.Registry))),
    )
    exampleClient := example.NewExampleService(conf.Conf.Micro.ProjectName, c)
    msgs := make([]*example.Msg, 0)
    msgs = append(msgs, &example.Msg{Name: "name", Password: "pw"})

    rep, err := exampleClient.Login(context.TODO(), &example.LoginMsgReq{Msg: msgs})
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(rep)

}
`
