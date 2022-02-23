package protos

import (
	"context"
	"io"
	"net/url"
	"plugin"

	"github.com/shavac/mp1p/cfg"
)

var (
	paMap = make(map[string]protoAdaptFunc)
)

type protoAdaptFunc func(string, io.Reader, *url.URL, ...[]byte) protoAdaptor

type protoAdaptor interface {
	GetResp() [][]byte
	Neg(context.Context) int //return -1 if not match, otherwise return matching position
	Handover(context.Context)
}

func RegisterProtoAdaptFunc(name string, f protoAdaptFunc) {
	paMap[name] = f
}

func GetProtoAdaptBuilder(name string) (protoAdaptFunc, error) {
	f, ok := paMap[name]
	if !ok {
		for _, path := range cfg.Config().Plugins.Paths {
			plugin.Open(path)
		}
	}
	return f, nil
}
