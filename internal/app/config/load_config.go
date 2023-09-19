package config

import (
	"github.com/spf13/viper"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type ConfigApp struct {
	App      App
	Database Database
	Log      logrusx.Config
}

type App struct {
	Debug bool   `mapstructure:"debug"`
	Port  string `mapstructure:"port"`
}

type Database struct {
	HostName     string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"database_name"`
}

func LoadConfig() (config ConfigApp, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}
