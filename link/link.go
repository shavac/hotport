package link

import (
	"io"
	"net"
	"sync"

	"go.uber.org/multierr"
)

var (
	allLinks = links{
		mu:    sync.RWMutex{},
		links: make(map[net.Addr]*Link),
	}
)

type links struct {
	mu    sync.RWMutex
	links map[net.Addr]*Link
}

func (l *links) regLink(lnk *Link) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.links[lnk.RemoteAddr] = lnk
}

func (l *links) unregLink(addr net.Addr) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.links, addr)
}

type Link struct {
	ServiceName             string
	RemoteAddr              net.Addr
	DialInConn, DialOutConn io.ReadWriteCloser
	TransportFunc           func()
}

func (l Link) Transport() {
	go l.TransportFunc()
}

func (l Link) Close() error {
	err := multierr.Append(l.DialInConn.Close(), l.DialOutConn.Close())
	return err
}

func RegisterLink(lnk *Link) {
	allLinks.regLink(lnk)
}

func UnRegisterLink(addr net.Addr) {
	allLinks.unregLink(addr)
}
