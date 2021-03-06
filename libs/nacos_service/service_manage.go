package nacos_service

import (
	"encoding/json"
	"fmt"
	"robot-api/libs/cache"
	"robot-api/libs/config"
)

func GetNacosRegisterService(conf *config.Config) {
	services := conf.DependentService
	data, _ := json.Marshal(&services)
	servicesMap := make(map[string]interface{})
	json.Unmarshal(data, &servicesMap)
	ss := servicesMap["AreaService:map"]
	fmt.Println(ss)
	// for key, _ := range servicesMap {
	// 	instance, _ := nacos_helper.SelectServiceInstance(nil, &ncopts.SelectServiceInstanceOptions{
	// 		ServiceName: name,
	// 		GroupName:   "DEFAULT_GROUP",
	// 		//GroupName:   os.Getenv("APP_GROUP_NAME"), // 区分不同运行环境，如 pre-production、production 而不需要更改配置，只需要运行时设置不同的环境变量
	// 	})
	// 	cache.Set(name, fmt.Sprintf("http://%s:%d", instance.Ip, instance.Port))
	// }
}

func GetServiceUrl(serviceName string) string {
	return cache.Get(serviceName)
}
