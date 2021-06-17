package config

import (
	"github.com/zhangsq-ax/service_config"
)

type Config struct {
	Self struct {
		Names   []string `yaml:"names"`   //服务名列表
		Name    string   `yaml:"name"`    // 服务名称
		Port    int      `yaml:"port"`    // 监听端口
		LogsDir string   `yaml:"logsDir"` // 接口查询存储目录
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

	return cfg.(*Config), nil
}
