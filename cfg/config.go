package cfg

import (
	"github.com/spf13/viper"
)

//type portConfigs = map[string]portConfig
//type protocolConfigs = map[string]portConfig

type config struct {
	Listen map[string]listenConfig
	Proto  map[string]protoConfig
}

type listenConfig struct {
	Addr     string
	Services []string
}

type protoConfig struct {
	Protocol  string
	ForwardTo string `mapstructure:"forward_to"`
	LoginUser string `mapstructure:"login_user"`
}

func ReadFromTomlFile(fname string) (*config, error) {
	cfg := &config{}
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigFile(fname)
	v.ReadInConfig()
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
