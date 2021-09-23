package template

var Handler = `package handler

import (
	"context"
	"encoding/json"
	"{{.Name}}/domain/model"
	"{{.Name}}/domain/service"
	"{{.Name}}/proto/example"
)

type Example struct {
	ExampleService service.IExampleService
}

func (e *Example) Hello(ctx context.Context,request *example.HelloReq, respose *example.HelloRep) (err error) {
	example := &model.Example{}
	jsonRequest,err := json.Marshal(request)
	if err != nil {
		return err
	}
	json.Unmarshal(jsonRequest, example)
	respose.Msg, err = e.ExampleService.Hello(example)
	return nil
}`
