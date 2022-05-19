package log

import (
	"github.com/shavac/mp1p/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l, _    = zap.NewProduction()
	cfg     = zap.NewProductionConfig()
	Fatalln = S().Panic
	Errorln = S().Error
	Warnln  = S().Warn
	Infoln  = S().Info
	Debugln = S().Debug
)

func L() *zap.Logger {
	return l
}

func S() *zap.SugaredLogger {
	return l.Sugar()
}

/*
func SetLogFileName(lf string) {
	cfg.OutputPaths = []string{path + "/" + lf}
	l, _ = cfg.Build()
}
*/

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

func Setup() {
	var err error
	conf := global.GetConfig().LogConfig
	switch conf.Level {
	case "DEBUG":
		cfg.Development = true
		println(111)
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
	if len(conf.LogPath) != 0 {
		cfg.OutputPaths = []string{conf.LogPath}
	}
	nl, err := cfg.Build()
	if err != nil {
		Fatalln(err)
	}
	l = nl
}
