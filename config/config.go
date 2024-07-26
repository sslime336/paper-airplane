package config

import (
	"log"

	"github.com/spf13/viper"
)

var appConfig AppConfig

type AppConfig struct {
	Bot BotConf `yaml:"bot"`
	Log LogConf `yaml:"log"`
}

type BotConf struct {
	Token string `yaml:"token"`
	Uin   uint64 `yaml:"uin"`
	AppId uint64 `yaml:"appId"`
}

type LogConf struct {
	Path string `yaml:"path"`
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read in config failed:", err)
	}
	viper.Unmarshal(&appConfig)
}

func Bot() BotConf {
	return appConfig.Bot
}

func Log() LogConf {
	return appConfig.Log
}

func App() AppConfig {
	return appConfig
}
