package log

import (
	"strings"

	"github.com/shavac/hotport/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	cfg     = zap.NewProductionConfig()
	l, _    = zap.NewProduction()
	s       = l.Sugar()
	Fatalln = s.Panic
	Errorln = s.Error
	Warnln  = s.Warn
	Infoln  = s.Info
	Debugln = s.Debug
)

func buildLoggers(cfg zap.Config) error {
	if nl, err := cfg.Build(); err != nil {
		return nil
	} else {
		l = nl
	}
	s = l.Sugar()
	Fatalln = s.Panic
	Errorln = s.Error
	Warnln = s.Warn
	Infoln = s.Info
	Debugln = s.Debug
	return nil
}

/*
func SetLogFileName(lf string) {
	cfg.OutputPaths = []string{path + "/" + lf}
	l, _ = cfg.Build()
}
*/

func SetLevel(lvl zapcore.Level) error {
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	return buildLoggers(cfg)
}

func SetStackTrace(trace bool) error {
	cfg.DisableStacktrace = !trace
	return buildLoggers(cfg)
}

func SetEncoding(enc string) error {
	cfg.Encoding = enc
	return buildLoggers(cfg)
}

func Setup() error {
	var err error
	conf := global.GetConfig().Log
	if strings.ToUpper(conf.Level) == "DEBUG" {
		cfg = zap.NewDevelopmentConfig()
	} else if cfg.Level, err = zap.ParseAtomicLevel(conf.Level); err != nil {
		cfg.Level = zap.NewAtomicLevel()
	}
	if conf.Format != "" {
		cfg.Encoding = conf.Format
	}
	if len(conf.LogPath) != 0 {
		cfg.OutputPaths = []string{conf.LogPath + "/hotport.log"}
	}
	return buildLoggers(cfg)
}
