package template

var Service = `package service

import (
	"{{.Name}}/domain/model"
)

type IExampleService interface {
	Hello(*model.Example) (string, error)
}

type ExampleHello struct {}


func NewExampleApi() *ExampleHello {
	return &ExampleHello{}
}

func (u *ExampleHello) Hello(example *model.Example) (string, error) {
	//业务
	hello := "Hello " + example.Name

	return hello,nil
}

`
