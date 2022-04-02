package cfg

import (
	"log"

	"github.com/shavac/mp1p/global"
	"github.com/spf13/viper"
)

func ReadFromPath(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("log")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	viper.SetConfigName("port")
	if err := viper.MergeInConfig(); err != nil {
		log.Println(err)
	}
	viper.SetConfigName("service")
	if err := viper.MergeInConfig(); err != nil {
		log.Println(err)
	}
	viper.SetConfigName("plugin")
	if err := viper.MergeInConfig(); err != nil {
		log.Println(err)
	}
	return viper.Unmarshal(global.Config)
}
