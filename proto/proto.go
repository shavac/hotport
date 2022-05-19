package proto

import (
	"context"
	"net"
	"net/url"
	"plugin"

	"github.com/shavac/hotport/errs"
	"github.com/shavac/hotport/link"
)

var (
	protoSvcMap = make(map[string]protoServiceFunc)
	pluginPath  = make(map[string]bool)
)

type protoServiceFunc func(string, *url.URL, ...string) (ProtoService, error)

type ProtoService interface {
	TryConn(context.Context, []byte, net.Conn) (*link.Link, []byte, bool) //link, bytes read in, match ok
	LocalURL() *url.URL
}

func RegisterProtoServiceFunc(name string, f protoServiceFunc) {
	protoSvcMap[name] = f
}

func AddPluginPath(path string) {
	pluginPath[path] = false
}

func GetProtoServiceBuilder(name string) (protoServiceFunc, error) {
	f, ok := protoSvcMap[name]
	if !ok {
		for path := range pluginPath {
			plugin.Open(path)
		}
		if f, ok = protoSvcMap[name]; !ok {
			return nil, errs.ErrNotImplemented
		}
	}
	return f, nil
}

func NewService(svcName, protoName string, tgtUrl *url.URL, args ...string) (ProtoService, error) {
	f, err := GetProtoServiceBuilder(protoName)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, errs.ErrNotImplemented
	}
	return f(svcName, tgtUrl, args...)
}
