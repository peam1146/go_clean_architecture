package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type App struct {
	AppName string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Debug   bool   `mapstructure:"debug"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSL      string `mapstructure:"ssl"`
}

type AppConfig struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
}

func LoadAppConfig() (config *AppConfig, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := AppConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf, err
}
