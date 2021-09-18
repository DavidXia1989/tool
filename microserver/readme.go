package template
var Readme = `### 环境要求

golang >= 1.13

` + "`" +  "`"  + "`" +`shell
export GO111MODULE=on
export GOPROXY=https://goproxy.cn/,direct
export GOPRIVATE=code.zm.shzhanmeng.com
{{.GoPath}}/bin 加入环境变量
GoPath/bin 加入环境变量
` + "`" +  "`"  + "`" +`

安装protoc

` + "`" +  "`"  + "`" +`
go get github.com/protocolbuffers/protobuf/v2
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro/v2
` + "`" +  "`"  + "`" +`

### 运行

生成protobuf文件

` + "`" +  "`"  + "`" +`
protoc -I. --go_out=. --micro_out=. proto/example.proto
` + "`" +  "`"  + "`" +`

启动

` + "`" +  "`"  + "`" +`
go mod tidy
go run main.go
` + "`" +  "`"  + "`" +`

`
