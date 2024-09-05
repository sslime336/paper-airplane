package config

import (
	"log"

	"github.com/spf13/viper"
)

func ParseConfig[T any](configUrl string) T {
	viper.SetConfigFile(configUrl)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read in config failed:", err)
	}
	var conf T
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal("unmarshal failed:", err)
	}
	return conf
}
