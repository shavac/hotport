package ports

import (
	"net"
	"net/url"

	p "github.com/shavac/mp1p/protos"
)

type Port struct {
	Name       string
	ListenAddr *net.TCPAddr
	Services   []p.ProtoAdaptor
}

func NewPort(name string, addr *net.TCPAddr) (*Port, error) {
	p := Port{Name: name, ListenAddr: addr}
	return &p, nil
}

func (p *Port) AddService(name, proto string, target *url.URL, args []byte) error {
	return nil
}
