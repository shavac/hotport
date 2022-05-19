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
	viper.SetConfigName("ports")
	if err := viper.MergeInConfig(); err != nil {
		log.Println(err)
	}
	viper.SetConfigName("services")
	if err := viper.MergeInConfig(); err != nil {
		log.Println(err)
	}
	viper.SetConfigName("plugins")
	if err := viper.MergeInConfig(); err != nil {
		log.Println(err)
	}
	return viper.Unmarshal(global.GetConfig())
}
