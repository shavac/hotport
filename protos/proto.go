package protos

import (
	"context"
	"io"
	"net/url"
	"plugin"
)

var (
	paMap      = make(map[string]protoAdaptFunc)
	pluginPath = make(map[string]bool)
)

type protoAdaptFunc func(string, io.Reader, *url.URL, ...[]byte) ProtoAdaptor

type ProtoAdaptor interface {
	GetResp() [][]byte
	Neg(context.Context, []byte) (int, bool) //offset, ok
	Handover(context.Context)
}

func RegisterProtoAdaptFunc(name string, f protoAdaptFunc) {
	paMap[name] = f
}

func AddPluginPath(path string) {
	pluginPath[path] = false
}

func GetProtoAdaptBuilder(name string) (protoAdaptFunc, error) {
	f, ok := paMap[name]
	if !ok {
		for path, _ := range pluginPath {
			plugin.Open(path)
		}
	}
	return f, nil
}
