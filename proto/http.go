package proto

import (
	"context"
	"net"
	"net/url"
	"regexp"
	"time"

	"github.com/shavac/hotport/log"
)

func init() {
	RegisterProtoServiceFunc("http", newHttpProt)
}

type httpProt struct {
	name   string
	fwdURL *url.URL
	reqs   []string
}

func newHttpProt(name string, furl *url.URL, reqs ...string) (ProtoService, error) {
	p := httpProt{
		name:   name,
		fwdURL: furl,
		reqs:   reqs,
	}
	return &p, nil
}

func (p httpProt) Name() string {
	return p.name
}

func (p httpProt) LocalURL() *url.URL {
	return p.fwdURL
}

func (p *httpProt) TryConn(ctx context.Context, msg NegMsg, in net.Conn) (net.Conn, NegMsg, bool) {
	in.SetReadDeadline(time.Now().Add(1 * time.Second))
	defer in.SetReadDeadline(time.Unix(0, 0))
	inb := []byte{}
	n, err := in.Read(inb)
	msg.inb = append(msg.inb, inb...)
	log.Debugln(in.RemoteAddr(), inb, n, err)
	ok, err := regexp.Match("^GET|POST|PUT|DELETE|TRACE|HEAD|OPTIONS|CONNECT|TRACE", msg.inb)
	if !ok || err != nil {
		return nil, msg, ok
	}
	hostPort := net.JoinHostPort(p.fwdURL.Hostname(), p.fwdURL.Port())
	if p.fwdURL.Port() == "" {
		hostPort = net.JoinHostPort(p.fwdURL.Hostname(), p.fwdURL.Scheme)
	}
	out, err := net.Dial("net", hostPort)
	if err != nil {
		log.Errorln("Connecting to ", hostPort, "ERROR: ", err)
		return nil, msg, false
	}
	log.Debugln("Connected to ", hostPort)
	return out, msg, ok
}

func (p httpProt) Transport(msg NegMsg, in, out net.Conn) {
	out.Write(msg.inb)
	in.Write(msg.outb)
	pipe(in, out)
}
