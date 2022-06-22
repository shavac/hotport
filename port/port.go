package port

import (
	"context"
	"net"
	"strings"
	"sync"

	"github.com/shavac/hotport/link"
	"github.com/shavac/hotport/log"
	"github.com/shavac/hotport/proto"
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

func (p *Port) String() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	svcStrs := []string{}
	for _, svc := range p.svcs {
		svcStrs = append(svcStrs, svc.Name())
	}
	s := p.lis.Addr().String() + " {" + strings.Join(svcStrs, ", ") + "}"
	return s
}

func (p *Port) Name() string {
	return p.name
}

func (p *Port) AddService(svcs ...proto.ProtoService) error {
	p.mu.Lock()
	defer p.mu.Unlock()
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
	p.mu.Lock()
	defer p.mu.Unlock()
	for {
		conn, err := p.lis.Accept()

		if err != nil {
			log.Errorln(err)
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
