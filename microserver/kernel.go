package template

var Kernel = `package kernel
import (
	"code.zm.shzhanmeng.com/go-common/mysql_xorm"
	"code.zm.shzhanmeng.com/go-common/redis"
	"code.zm.shzhanmeng.com/go-common/logging"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"context"
	"os"
	"os/signal"
	"syscall"
	"errors"
	"{{.Name}}/common"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"go.uber.org/zap"
)

type server struct {
	ProjectName string	` + "`" + `yaml:"project_name"` + "`" + `
	RunMode 	string	` + "`" + `yaml:"run_mode"` + "`" + `
	HttpPort	string	` + "`" + `yaml:"http_port"` + "`" + `
	GrpcPort	string	` + "`" + `yaml:"grpc_port"` + "`" + `
	Registry	string	` + "`" + `yaml:"registry"` + "`" + `
	Mysql		[]mysql_xorm.XmsyqlConf	` + "`" + `yaml:"mysql"` + "`" + `
	Redis		[]redis.RedisConf			` + "`" + `yaml:"redis"` + "`" + `
	Log			logging.LogConf			` + "`" + `yaml:"log"` + "`" + `
}
// 系统配置
var ServerSetting = &server{}

// 配置句柄
var ConfigContent []byte
var MicroServer *micro.Service

// 动态映射获取配置
func GetConfig(conf interface{}) error {
	err := yaml.Unmarshal(ConfigContent, conf)
	if err != nil {
		return errors.New("读取配置文件失败 :" + err.Error())
	}
	return nil
}

func init(){
	setupConf() //  读取文件配置
	setupLog() // 初始化日志
}

//监听并处理信号
func SetSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-c:
				cancel()
				return
			}
		}
	}()
	return ctx
}
// 读取默认配置
func setupConf(){
	filePath, err := filepath.Abs(filepath.Join("conf", "app.yaml"))
	if err != nil {
		panic(errors.New("setting.Setup, fail to parse 'conf/app.yaml'" + err.Error()))
	}
	ConfigContent, err = ioutil.ReadFile(filePath)
	if err != nil {
		panic(errors.New("setting.Setup parse app.yaml err:'" + err.Error()))
	}
	err = GetConfig(ServerSetting)
	if  err != nil{
		panic(err)
	}
}


// 初始化日志
func setupLog(){
	// 设置日志开发模式
	ServerSetting.Log.Debug = ServerSetting.RunMode
	logging.InitLogger(ServerSetting.Log)
}

// 初始化mysql
func SetupMysql(){
	_, err :=mysql_xorm.NewClients(ServerSetting.Mysql)
	if err != nil {
		logging.ZapLogger.Warn("链接mysql数据库失败")
		panic(err)
	}
}

// 初始化redis
func SetupRedis(){
	if common.IsStructureEmpty(redis.RedisConf{},ServerSetting.Redis) {
		return
	}
	err := redis.NewClients(ServerSetting.Redis)
	if err != nil {
		panic(errors.New("link redis err:'" + err.Error()))
	}
}

// HTTP服务
func SetupHttp(r *gin.Engine){
	if ServerSetting.HttpPort != "" {
		err := r.Run("0.0.0.0" + ":" + ServerSetting.HttpPort)
		if err != nil {
			panic("server服务器启动失败")
		}
	}
}

//初始化grpc服务
func GrpcInit(){
	grpcAddr := "0.0.0.0:"+ServerSetting.GrpcPort

	ser := micro.NewService(
		micro.Name("zmtooltest"),
		micro.Address(grpcAddr),
		micro.Registry(etcd.NewRegistry(registry.Addrs(ServerSetting.Registry))),
		// 限流5
		micro.WrapHandler(ratelimit.NewHandlerWrapper(100)),
	)
	MicroServer = &ser
	// Init will parse the command line flags.
	(*MicroServer).Init()
}

func GrpcStart(){
	//启动并监听服务
	if err := (*MicroServer).Run(); err != nil {
		logging.ZapLogger.Info("micro grpc 启动失败",zap.Error(err))
		fmt.Println(err)
	}
}
`
