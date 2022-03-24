package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l, _ = zap.NewProduction()
	cfg  = zap.NewProductionConfig()
	path = "/tmp"
)

func L() *zap.Logger {
	return l
}

func S() *zap.SugaredLogger {
	return l.Sugar()
}

func SetLogPath(p string) {
	path = p
}

func SetLogFileName(lf string) {
	cfg.OutputPaths = []string{path + "/" + lf}
	l, _ = cfg.Build()
}

func SetLevel(lvl zapcore.Level) {
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	l, _ = cfg.Build()
}

func SetStackTrace(trace bool) {
	cfg.DisableStacktrace = !trace
	l, _ = cfg.Build()
}

func SetEncoding(enc string) {
	cfg.Encoding = enc
	l, _ = cfg.Build()
}

func Setup(conf *Config) {
	var err error
	if conf == nil {
		return
	}
	switch conf.Level {
	case "DEBUG":
		cfg.Development = true
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "INFO":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "WARN":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "ERROR":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "PANIC":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	}
	cfg.Encoding = conf.Encoding
	path = conf.LogPath
	cfg.OutputPaths = []string{conf.LogPath}
	nl, err := cfg.Build()
	if err != nil {
		l.Fatal(err.Error())
	}
	l = nl
}
