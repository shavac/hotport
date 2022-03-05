package ports

import (
	"net"

	"github.com/shavac/mp1p/protos"
)

type Port struct {
	Name       string
	ListenAddr *net.TCPAddr
	Services   []protos.ProtoAdaptor
}

func NewPort(name string, addr *net.TCPAddr) (*Port, error) {
	p := Port{Name: name, ListenAddr: addr}
	return &p, nil
}

func (p *Port) AddServiceByConfig(protoCfg protos.Config) error {
	return nil
}
