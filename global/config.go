package global

var (
	cfg = &config{
		PortConfig:    make(map[string]portConfig),
		ServiceConfig: make(map[string]serviceConfig),
		LogConfig: logConfig{
			LogPath:  "",
			Level:    "DEBUG",
			Encoding: "console",
		},
	}
)

func GetConfig() *config {
	return cfg
}

type config struct {
	PortConfig    map[string]portConfig
	ServiceConfig map[string]serviceConfig
	LogConfig     logConfig
	PluginConfig  pluginConfig
}

type portConfig struct {
	ListenAddr string `mapstructure:"addr"`
	Services   []string
}

type serviceConfig struct {
	Protocol     string
	ForwardToURL string   `mapstructure:"forward_to"`
	Arguments    []string `mapstructure:"arguments"`
}

type pluginConfig struct {
	Paths []string
}

type logConfig struct {
	LogPath  string
	Level    string
	Encoding string
}
