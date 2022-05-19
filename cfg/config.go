package cfg

import (
	"github.com/shavac/hotport/log"

	"github.com/shavac/hotport/global"
	"github.com/spf13/viper"
)

func ReadFromPath(path string) error {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("ports")
	if err := v.ReadInConfig(); err != nil {
		log.Errorln(err)
	}
	v.SetConfigName("log")
	if err := v.MergeInConfig(); err != nil {
		log.Errorln(err)
	}
	v.SetConfigName("services")
	if err := v.MergeInConfig(); err != nil {
		log.Errorln(err)
	}
	v.SetConfigName("plugins")
	if err := v.MergeInConfig(); err != nil {
		log.Errorln(err)
	}
	err := v.Unmarshal(global.GetConfig())
	return err
}
