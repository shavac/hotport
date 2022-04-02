package main

import (
	"net"
	"net/netip"
	"net/url"

	"github.com/shavac/mp1p/cfg"
	"github.com/shavac/mp1p/cmd"
	"github.com/shavac/mp1p/log"
	"github.com/shavac/mp1p/port"
	"github.com/shavac/mp1p/proto"
)

var (
	cfgChgEvt = make(chan bool)
)

func init() {
	cmd.Execute()
}

var allPorts = make(map[string]*port.Port)

func main() {
	for {
		//fmt.Println(cfg.Config())
		for portName, portCfg := range cfg.Config().Port {
			laddr, err := net.ResolveTCPAddr("tcp", portCfg.ListenAddr)
			if err != nil {
				log.S().Error(err)
				continue
			}
			p, err := port.NewPort(portName, laddr)
			if err != nil {
				log.S().Error(err)
				continue
			}
			for _, sName := range portCfg.Services {
				sCfg := cfg.Config().Service[sName]
				u, err := url.Parse(sCfg.ForwardToURL)
				if err != nil {
					log.S().Error(err)
					continue
				}
				s, err := proto.NewService(sName, sCfg.Protocol, u, sCfg.Arguments...)
				if err != nil {
					log.S().Error(err)
					continue
				}
				p.AddService(s)
			}
			allPorts[portName] = p
			p.Accept()
			defer func(pName string) {
				allPorts[pName].Close()
				delete(allPorts, pName)
			}(portName)
		}
		<-cfgChgEvt
	}
}

func nipAddrPortToAddr(nap netip.AddrPort) net.Addr {
	return nil
}
