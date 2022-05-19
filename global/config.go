package global

var (
	cfg = &config{
		Port:    make(map[string]portConfig),
		Service: make(map[string]serviceConfig),
		Log: logConfig{
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
	Port    map[string]portConfig
	Service map[string]serviceConfig
	Log     logConfig
	Plugin  pluginConfig
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
