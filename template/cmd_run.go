package template

var Cmd_run = `package common

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// 运行cmd命令标准模式
func CmdRunStandard(name string, arg ...string) (outStr, errStr string, err error) {
	//exec.Command("bash", "-c", "java -version 2>&1 | sed '1!d'|sed 's/\"//g' | awk '{print $NF}' | awk -F'.' '{print $1$2}'")
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err = cmd.Run()
	outStr, errStr = string(stdout.Bytes()), string(stderr.Bytes())
	outStr = strings.Trim(outStr, "\n")
	errStr = strings.Trim(errStr, "\n")

	return outStr, errStr, err
}

// 运行cmd命令
func CmdRun(name string, arg ...string)(string,error){
	outStr, errStr, err := CmdRunStandard(name,arg...)
	if err!=nil {
		return "",err
	}
	if errStr!="" {
		return errStr,errors.New(errStr)
	}
	return outStr,nil
}
`
