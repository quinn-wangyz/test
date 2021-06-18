package main

import (
	"flag"
	"fmt"
	"os"
	"robot-api/libs/config"
	"robot-api/libs/nacos_service"
	"robot-api/router"

	"github.com/golang/glog"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	nacos_helper "github.com/zhangsq-ax/nacos-helper-go"
	ncopts "github.com/zhangsq-ax/nacos-helper-go/options"
)

func init() {
	flag.Parse()
}

func main() {
	// 获取配置
	conf, err := config.GetConfig()
	if err != nil {
		glog.Fatalln(err)
	}

	//向 Nacos 注册服务
	err = registerServiceInstance(conf)
	fmt.Println(conf)
	if err != nil {
		glog.Fatalln(err)
	}

	//获取 nacos 服务
	nacos_service.GetNacosRegisterService(conf)

	// 启动 Web 服务接口
	err = startWeb(conf.Self.Port)
	if err != nil {
		glog.Fatalln(err)
	}
}

func registerServiceInstance(conf *config.Config) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	glog.Infof(">>>>> Register service instance: ServiceName - %s, IP - %s, Port - %d\n", conf.Self.Name, hostname, conf.Self.Port)
	return nacos_helper.RegisterServiceInstance(nil, &ncopts.RegisterServiceOptions{
		ServiceName: conf.Self.Name,
		Ip:          hostname,
		Port:        uint64(conf.Self.Port),
		GroupName:   "DEFAULT_GROUP",
		//GroupName:   os.Getenv("APP_GROUP_NAME"), // 区分不同运行环境，如 pre-production、production 而不需要更改配置，只需要运行时设置不同的环境变量
	})

}

func startWeb(port int) error {
	app := iris.New()
	app.Use(logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
	}))

	app.UseRouter(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	}))

	app.PartyFunc("/", router.Router)

	glog.Infof(">>>>> Start web service on port %d\n", port)
	// 启动服务
	return app.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutServerError(iris.ErrServerClosed))
}
