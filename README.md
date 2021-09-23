##环境要求

golang 版本 >=1.13

设置环境变量

```shell
#代理&私有库地址
export GOPRIVATE=code.zm.shzhanmeng.com
export GOPROXY=https://goproxy.cn
#设置gopath/bin路径
{{gopath}}/bin 加入环境变量
```

### 获取工具

```
go get code.zm.shzhanmeng.com/go-common/zmtool
```

### 使用

```
// 生成一个单体应用
zmtool example
// 生成一个服务
zmtool exampleServ server
// 生成一个api网关
zmttol exampleApi api

```


### 目录结构
```
├── example
│   ├── main.go //入口文件
│   ├── go-build.sh //构建脚本
│   └── go.mod //依赖管理配置文件
├── example\shell
│   ├── check_monitor.sh 
│   ├── monitor_exec.sh
│   └── publish-script.sh
├── example\conf
│   ├── app.test.yaml //配置文件
│   └── config.go
├── example\routers
│   └── route.go //路由
├── example\controllers\exampleController
│   └── example.go //控制器
├── example\kernel
│   └── kernel.go
└── example\common
├── cmd_run.go
├── common.go
└── response.go
```