package proto

import (
	"context"
	"errors"
	"io"
	"net"
	"net/url"
	"plugin"

	"github.com/shavac/hotport/errs"
	"github.com/shavac/hotport/log"
)

var (
	protoSvcMap = make(map[string]protoServiceFunc)
	pluginPath  = make(map[string]bool)
)

type NegMsg struct {
	inb, outb []byte
}

type protoServiceFunc func(string, *url.URL, ...string) (ProtoService, error)

type ProtoService interface {
	Name() string
	TryConn(context.Context, NegMsg, net.Conn) (net.Conn, NegMsg, bool) //link, bytes read in, match ok
	LocalURL() *url.URL
	Transport(NegMsg, net.Conn, net.Conn) // NegMsg, in, out
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

func pipe(in, out io.ReadWriteCloser) {
	inCh, outCh := readIntoChan(in), readIntoChan(out)
LOOP:
	for {
		select {
		case b1, ok := <-inCh:
			if !ok {
				out.Close()
				break LOOP
			}
			out.Write(b1)
		case b2, ok := <-outCh:
			if !ok {
				in.Close()
				break LOOP
			}
			in.Write(b2)
		}
	}
}

func readIntoChan(in io.ReadWriteCloser) chan []byte {
	inCh := make(chan []byte)
	go func() {
		defer func() {
			close(inCh)
		}()
		for {
			bs := make([]byte, 128)
			n, err := in.Read(bs)
			if n > 0 {
				inCh <- bs
			}
			if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
				break
			}
			if err != nil {
				log.Errorln("Reading from network error:, ", err)
				break
			}
		}
	}()
	return inCh
}
