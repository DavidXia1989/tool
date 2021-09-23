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

func (e *Example) Login(ctx context.Context,request *example.LoginMsgReq, respose *example.LoginMsgRes) (err error) {
	example := &model.Example{}
	jsonRequest,err := json.Marshal(request.Msg[0])
	if err != nil {
		return err
	}
	json.Unmarshal(jsonRequest, example)
	respose.Result, err = e.ExampleService.Login(example)
	return nil
}`
