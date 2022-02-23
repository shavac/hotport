package ports

import (
	"context"
	"net"
)

type Port struct {
	Name       string
	ListenAddr *net.TCPAddr
	Services   []*Service
}

type Service struct {
	Name      string
	MatchFunc func(context.Context) int
}
