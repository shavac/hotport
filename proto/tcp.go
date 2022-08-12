package proto

import (
	"context"
	"net"
	"net/url"

	"github.com/shavac/hotport/log"
)

func init() {
	RegisterProtoServiceFunc("tcp", newTcpProt)
	////////////////////////////////
	// RegisterProtoServiceFunc("ssh", newTcpProt)
	// RegisterProtoServiceFunc("rdp", newTcpProt)
	// RegisterProtoServiceFunc("https", newTcpProt)
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

func (p tcpProt) TryConn(ctx context.Context, msg NegMsg, in net.Conn) (net.Conn, NegMsg, bool) {
	//d := net.Dialer{}

	hostPort := net.JoinHostPort(p.fwdURL.Hostname(), p.fwdURL.Port())
	if p.fwdURL.Port() == "" {
		hostPort = net.JoinHostPort(p.fwdURL.Hostname(), p.fwdURL.Scheme)
	}
	out, err := net.Dial("tcp", hostPort)
	if err != nil {
		log.Errorln("Connectionting to ", hostPort, "ERROR: ", err)
		return nil, msg, false
	}
	log.Debugln("Connected to ", hostPort)

	return out, msg, true
}

func (p tcpProt) Transport(msg NegMsg, in, out net.Conn) {
	out.Write(msg.inb)
	in.Write(msg.outb)
	pipe(in, out)
}
