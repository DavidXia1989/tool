package template
var Sh_go_build = `#!/bin/bash
export GO111MODULE="on"
export GOPROXY="code.zm.shzhanmeng.com"
{{.GoRoot}}/go mod tidy
{{.GoRoot}}/go build -o {{.Name}} main.go`