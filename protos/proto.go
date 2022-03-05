package protos

import (
	"context"
	"io"
	"net/url"
	"plugin"
)

var (
	paMap = make(map[string]protoAdaptFunc)
)

type protoAdaptFunc func(string, io.Reader, *url.URL, ...[]byte) ProtoAdaptor

type ProtoAdaptor interface {
	GetResp() [][]byte
	Neg(context.Context) int //return -1 if not match, otherwise return matching position
	Handover(context.Context)
}

func RegisterProtoAdaptFunc(name string, f protoAdaptFunc) {
	paMap[name] = f
}

func GetProtoAdaptBuilder(name string, path []string) (protoAdaptFunc, error) {
	f, ok := paMap[name]
	if !ok {
		for _, path := range path {
			plugin.Open(path)
		}
	}
	return f, nil
}
