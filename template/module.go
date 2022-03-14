package template

var Module = `module {{.Name}}

go 1.14

require (
	github.com/DavidXia1989/cron v0.0.0-20201222054010-a8ea79fea7e1
	github.com/DavidXia1989/logging v0.0.0-20210207090756-8be76f7db7dc
	github.com/DavidXia1989/mysql_xorm v0.0.0-20210118111623-91ade31e7971
	github.com/DavidXia1989/redis v0.0.0-20210918022451-021f6ae5440e
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	go.uber.org/zap v1.17.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	xorm.io/core v0.7.3 // indirect
)
`
