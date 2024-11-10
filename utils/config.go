package utils

import (
	"log"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func init() {
	Viper = viper.New()
	Viper.AddConfigPath(".")
	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AutomaticEnv()

	err := Viper.ReadInConfig()

	if err != nil {
		log.Panic("[Cannot read config]: " + err.Error())
	}
}
