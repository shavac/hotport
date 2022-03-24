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

<<<<<<< HEAD
func NewPortFromConfig(cfg.)

func (p *Port) Accept() {

}

type Service struct {
	Name      string
	MatchFunc func(context.Context) int
=======
func (p *Port) AddService(name, proto string, target *url.URL, args []byte) error {
	return nil
>>>>>>> 7f88a793544650098c93da959b661526aed8f424
}
