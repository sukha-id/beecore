package config

import (
	"github.com/spf13/viper"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type ConfigApp struct {
	App      App            `mapstructure:"app"`
	Database Database       `mapstructure:"database"`
	Log      logrusx.Config `mapstructure:"log"`
}

type App struct {
	Debug   bool   `mapstructure:"debug"`
	Port    string `mapstructure:"port"`
	Timeout int    `mapstructure:"timeout"`
}

type Database struct {
	HostName              string `mapstructure:"hostname"`
	Port                  string `mapstructure:"port"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	DatabaseName          string `mapstructure:"database_name"`
	MaxIdleConnection     int    `mapstructure:"max_idle_connection"`
	MaxOpenConnection     int    `mapstructure:"max_open_connection"`
	MaxLifetimeConnection int    `mapstructure:"max_lifetime_connection"`
}

func LoadConfig(pathFile string) (config ConfigApp, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(pathFile)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}
