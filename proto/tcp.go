package proto

import (
	"context"
	"io"
	"net"
	"net/url"

	"github.com/shavac/hotport/link"
	"github.com/shavac/hotport/log"
)

func init() {
	RegisterProtoServiceFunc("tcp", newTcpProt)
	////////////////////////////////
	RegisterProtoServiceFunc("ssh", newTcpProt)
	RegisterProtoServiceFunc("rdp", newTcpProt)
	RegisterProtoServiceFunc("http", newTcpProt)
	RegisterProtoServiceFunc("https", newTcpProt)
}

type tcpProt struct {
	name   string
	fwdURL *url.URL
	reqs   []string
}

func newTcpProt(name string, furl *url.URL, reqs ...string) (ProtoService, error) {
	p := tcpProt{
		name:   name,
		fwdURL: furl,
		reqs:   reqs,
	}
	return &p, nil
}

func (p tcpProt) Name() string {
	return p.name
}

func (p tcpProt) LocalURL() *url.URL {
	return p.fwdURL
}

func (p tcpProt) TryConn(ctx context.Context, b []byte, in net.Conn) (*link.Link, []byte, bool) {
	//d := net.Dialer{}

	hostPort := net.JoinHostPort(p.fwdURL.Hostname(), p.fwdURL.Port())
	if p.fwdURL.Port() == "" {
		hostPort = net.JoinHostPort(p.fwdURL.Hostname(), p.fwdURL.Scheme)
	}
	log.S().Error(hostPort)
	out, err := net.Dial("tcp", hostPort)
	if err != nil {
		log.S().Error(err)
		return nil, []byte{}, false
	}
	l := link.Link{
		ServiceName: p.name,
		RemoteAddr:  in.RemoteAddr(),
		DialInConn:  in,
		DialOutConn: out,
		TransportFunc: func() {
			pipe(in, out)
		},
	}
	go l.Transport()
	return &l, []byte{}, true
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
		for {
			bs := make([]byte, 128)
			n, err := in.Read(bs)
			if n > 0 {
				inCh <- bs
			}
			if err != nil {
				close(inCh)
				if err != io.EOF {
					log.S().Error(err)
				}
				break
			}
		}
	}()
	return inCh
}
