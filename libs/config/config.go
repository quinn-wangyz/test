package config

import (
	"github.com/zhangsq-ax/service_config"
)

var Conf *Config

type Config struct {
	DependentService struct {
		MapService struct {
			ServiceName string `yaml:"serviceName"`
			Group       string `yaml:"group"`
			NameSpaceId string `yaml:"nameSpaceId"`
		} `yaml:"mapService"`
		RobotService struct {
			ServiceName string `yaml:"serviceName"`
			Group       string `yaml:"group"`
			NameSpaceId string `yaml:"nameSpaceId"`
		} `yaml:"robotService"`
		ElevatorService struct {
			ServiceName string `yaml:"serviceName"`
			Group       string `yaml:"group"`
			NameSpaceId string `yaml:"nameSpaceId"`
		} `yaml:"elevatorService"`
		BusinessService struct {
			ServiceName string `yaml:"serviceName"`
			Group       string `yaml:"group"`
			NameSpaceId string `yaml:"nameSpaceId"`
		} `yaml:"businessService"`
		AreaService struct {
			ServiceName string `yaml:"serviceName"`
			Group       string `yaml:"group"`
			NameSpaceId string `yaml:"nameSpaceId"`
		} `yaml:"areaService"`
	} `yaml:"dependentService"`
	Self struct {
		Name    string `yaml:"name"`    //注册服务名称
		Port    int    `yaml:"port"`    // 监听端口
		LogsDir string `yaml:"logsDir"` // 接口查询存储目录
	} `yaml:"self"`
}

func newConfig() interface{} {
	return &Config{}
}

func GetConfig() (*Config, error) {
	provider, err := service_config.GetConfigProvider(service_config.NewConfigProviderOptions(service_config.ConfigFormat_YAML, newConfig), false)
	if err != nil {
		return nil, err
	}
	cfg, err := provider.Config()
	if err != nil {
		return nil, err
	}
	Conf = cfg.(*Config)
	return Conf, nil
}
