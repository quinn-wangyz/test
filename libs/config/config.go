package config

import (
	"github.com/zhangsq-ax/service_config"
)

var Conf *Config

type Config struct {
	Self struct {
		Name                string   `yaml:"name"`                //注册服务名称
		Names               []string `yaml:"names"`               //服务名列表
		MapSerivceName      string   `yaml:"mapSerivceName"`      // 地图服务名称
		RobotSerivceName    string   `yaml:"robotSerivceName"`    // 机器人服务名称
		ElevatorSerivceName string   `yaml:"elevatorSerivceName"` // 电梯服务名称
		Port                int      `yaml:"port"`                // 监听端口
		LogsDir             string   `yaml:"logsDir"`             // 接口查询存储目录
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
