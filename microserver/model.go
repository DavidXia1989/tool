package template

var Model = `package model

type Example struct {
	Id   int64 ` + "`" + `json:"id""` + "`" + `
	Name string ` + "`" + `json:"name""` + "`" + `
	Password  string ` + "`" + `json:"password""` + "`" + `
}
`