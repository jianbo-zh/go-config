package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// viperInstance viperInstance 库实例
var viperInstance *viper.Viper

type ConfigOption struct {
	Type      string
	File      string
	Name      string
	Paths     []string
	EnvPrefix string
}

func Setup(option ConfigOption) error {
	if viperInstance != nil {
		return nil
	}

	viperInstance = viper.New()

	if option.Type != "" {
		viperInstance.SetConfigType(option.Type)
	}

	if option.File != "" {
		viperInstance.SetConfigFile(option.File)
	}

	if option.Name != "" {
		viperInstance.SetConfigName(option.Name)
	}

	for _, path := range option.Paths {
		if path != "" {
			viperInstance.AddConfigPath(path)
		}
	}

	err := viperInstance.ReadInConfig()
	if err != nil {
		return fmt.Errorf("viper read in config error: %w", err)
	}

	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	viperInstance.SetEnvPrefix(option.EnvPrefix)

	// viperInstance.Get() 时，优先读取环境变量
	viperInstance.AutomaticEnv()

	return nil
}
