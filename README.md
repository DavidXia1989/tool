设置环境变量
export GOPRIVATE=code.zm.shzhanmeng.com
export GOPROXY=https://goproxy.cn

golang 版本 >=1.13

go get code.zm.shzhanmeng.com/go-common/zmtool

{{gopath}}/bin 加入环境变量

zmtool example   会在当前目录下创建一个 example目录

zmtool example

Creating service  in

.
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