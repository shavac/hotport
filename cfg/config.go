package cfg

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	cfg      = &config{}
	onCfgChg = []func(){}
)

type config struct {
	Plugins pluginConfig
	Port    map[string]portConfig
	Service map[string]serviceConfig
}

type pluginConfig struct {
	Paths []string
}

type portConfig struct {
	ListenAddr string
	Services   []string
}

type serviceConfig struct {
	Protocol     string
	ForwardToURL string `mapstructure:"forward_to"`
	LoginUser    string `mapstructure:"login_user"`
}

func OnChange(f func()) {
	onCfgChg = append(onCfgChg, f)
}

func Config() config {
	return *cfg
}

func ReadFromPath(path string, typ string) error {
	switch typ {
	case "net":
	default:
		viper.SetConfigName("mp1p")
		viper.AddConfigPath(path)
		//viper.SetConfigType(typ)
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			if in.Op != fsnotify.Write {
				return
			}
			if err := viper.ReadInConfig(); err != nil {
				return
			}
			if err := viper.Unmarshal(cfg); err != nil {
				return
			}
			for _, f := range onCfgChg {
				go f()
			}
		})
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	return viper.Unmarshal(cfg)
}

func ReadFromFile(fname, typ string) error {
	switch typ {
	case "net":
	default:
		viper.SetConfigType(typ)
		viper.SetConfigFile(fname)
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			if in.Op != fsnotify.Write {
				return
			}
			if err := viper.ReadInConfig(); err != nil {
				return
			}
			if err := viper.Unmarshal(cfg); err != nil {
				return
			}
			for _, f := range onCfgChg {
				go f()
			}
		})
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	return viper.Unmarshal(cfg)
}
