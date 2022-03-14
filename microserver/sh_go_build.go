package template
var Sh_go_build = `#!/bin/bash
export GO111MODULE="on"
//export GOPRIVATE=code.zm.shzhanmeng.com
export GOPROXY=https://goproxy.cn
{{.GoRoot}}/go mod tidy
{{.GoRoot}}/go build -o {{.Name}} main.go`