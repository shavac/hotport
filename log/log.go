package log

import (
	"strings"

	"github.com/shavac/hotport/global"
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
	conf := global.GetConfig().Log
	if strings.ToUpper(conf.Level) == "DEBUG" {
		cfg = zap.NewDevelopmentConfig()
	} else if cfg.Level, err = zap.ParseAtomicLevel(conf.Level); err != nil {
		cfg.Level = zap.NewAtomicLevel()
	}
	cfg.Encoding = conf.Format
	if len(conf.LogPath) != 0 {
		cfg.OutputPaths = []string{conf.LogPath + "/hotport.log"}
	}
	nl, err := cfg.Build()
	if err != nil {
		Fatalln(err)
	}
	l = nl
}
