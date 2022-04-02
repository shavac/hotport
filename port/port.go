package port

import (
	"context"
	"net"
	"sync"

	"github.com/shavac/mp1p/link"
	"github.com/shavac/mp1p/log"
	"github.com/shavac/mp1p/proto"
)

type Port struct {
	mu   sync.RWMutex
	name string
	lis  *net.TCPListener
	svcs []proto.ProtoService
}

func NewPort(name string, addr *net.TCPAddr) (*Port, error) {
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}
	p := Port{name: name, lis: lis}
	return &p, nil
}

func (p *Port) Name() string {
	return p.name
}

func (p *Port) AddService(svcs ...proto.ProtoService) error {
	p.svcs = append(p.svcs, svcs...)
	return nil
}

func (p *Port) Accept() {
	go p.accept()
}

func (p *Port) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.lis.Close()
}

func (p *Port) accept() {
	for {
		conn, err := p.lis.Accept()

		if err != nil {
			log.S().Error(err)
			continue
		}
		//conn.SetDeadline(time.Now().Add((10 * time.Second)))
		go func() {
			var l *link.Link
			msg, ok := []byte{}, false
			for _, pa := range p.svcs {
				//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				ctx := context.Background()
				l, msg, ok = pa.TryConn(ctx, msg, conn)
				if ok {
					link.RegisterLink(l)
					break
				} else {
					conn.Close()
				}
			}
		}()
	}
}
