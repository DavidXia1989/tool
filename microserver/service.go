package template

var Service = `package service

import (
	"code.zm.shzhanmeng.com/go-common/mysql_xorm"
	"code.zm.shzhanmeng.com/go-common/logging"
	"fmt"
	"go.uber.org/zap"
	"{{.Name}}/domain/model"
)

type IExampleService interface {
	Login(*model.Example) (int64, error)
}

type ExampleLogin struct {}


func NewExampleApi() *ExampleLogin {
	return &ExampleLogin{}
}

func (u *ExampleLogin) Login(example *model.Example) (int64, error) {
	//业务
	
	return id,nil
}
`
