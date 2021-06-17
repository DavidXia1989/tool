package template

var Response = `package common

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Meta interface{} ` + "`" + `json:"meta"` + "`" +`
	Code int         ` + "`" + `json:"error_code"` + "`" +`
	Msg  string      ` + "`" + `json:"error_message"` + "`" +`
	Data interface{} ` + "`" + `json:"data"` + "`" +`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, errorMsg string, data interface{}, meta interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  errorMsg,
		Data: data,
		Meta: meta,
	})
	return
}

func (g *Gin) ResoponseSucess(data,meta interface{}){
	g.Response(200,0,"",data, meta)
	return
}

func (g *Gin) ResponnseFailure(error_code int, error_msg string){
	g.Response(200,error_code, error_msg,make(map[string]interface{}), make(map[string]interface{}))
	return
}

func (this *Gin) InputString(key string, defs ...string) string {
	def := ""
	if len(defs) > 0 {
		def = defs[0]
	}
	val := this.C.Query(key)
	if val == "" {
		val = this.C.PostForm(key)
	}
	if val == "" {
		return def
	}
	return val
}

func (this *Gin) InputInt(key string, defs ...int) int {
	def := 0
	if len(defs) > 0 {
		def = defs[0]
	}
	val := this.C.Query(key)
	if val == "" {
		val = this.C.PostForm(key)
	}
	if val == "" {
		return def
	}
	t,_ := strconv.Atoi(val)

	return t
}`
