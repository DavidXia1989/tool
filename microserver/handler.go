package template

var Handler = `package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"{{.Name}}/domain/model"
	"{{.Name}}/domain/service"
	"{{.Name}}/proto/example"
)

type Example struct {
	ExampleService service.IExampleService
}

func (e *Example) Login(ctx context.Context,request *example.LoginMsgReq, respose *example.LoginMsgRes) (err error) {
	example := &model.Example{}
	jsonRequest,err := json.Marshal(request)
	if err != nil {
		return err
	}
	json.Unmarshal(jsonRequest, example)
	fmt.Println(user)
	respose.UserId, err = e.ExampleService.Login(example)
	return nil
}`
