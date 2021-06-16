package nacos_service

import (
	"fmt"
	"robot-api/libs/cache"
	"robot-api/libs/config"

	nacos_helper "github.com/zhangsq-ax/nacos-helper-go"
	ncopts "github.com/zhangsq-ax/nacos-helper-go/options"
)

func InitService(conf *config.Config) {
	instance, err := nacos_helper.SelectServiceInstance(nil, &ncopts.SelectServiceInstanceOptions{
		ServiceName: conf.Self.Name,
		GroupName:   "DEFAULT_GROUP",
		//GroupName:   os.Getenv("APP_GROUP_NAME"), // 区分不同运行环境，如 pre-production、production 而不需要更改配置，只需要运行时设置不同的环境变量
	})
	fmt.Println(instance.Ip)
	fmt.Println(instance.Port)
	if err != nil {

	}
}

func ServiceManage() {
	cache.Set("sss", "ssss")
}

func GetManage() string {
	return cache.Get("sss")
}
