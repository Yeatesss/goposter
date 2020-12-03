package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConfigPath string
	ConfigName string
}

func NewConfig(path, name string) *Config {
	return &Config{
		ConfigPath: path,
		ConfigName: name,
	}
}

// 读取配置
func (c *Config) InitConfig() error {

	viper.AddConfigPath(c.ConfigPath)
	viper.SetConfigName(c.ConfigName)
	viper.SetConfigType("yaml")

	// 从环境变量总读取
	//viper.AutomaticEnv()
	//viper.SetEnvPrefix("poster")
	//viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	return viper.ReadInConfig()
}
