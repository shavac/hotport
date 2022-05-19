package port

type Config struct {
	ListenAddr string `mapstructure:"addr"`
	Services   []string
}
