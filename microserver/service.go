package template

var Service = `package service

import (
	"code.zm.shzhanmeng.com/go-common/mysql_xorm"
	"code.zm.shzhanmeng.com/go-common/logging"
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
	var res model.Example
	err := mysql_xorm.GetMysqlClient("gomicro").Table("user").Where("name",example.Name).Find(&res)
	if err != nil {
		logging.ZapLogger.Error("注册失败",zap.String("errmsg", err.Error()))
		return 0,err
	}

	return res.Id,nil
}
`
