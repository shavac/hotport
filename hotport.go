package main

import (
	"fmt"
	"net"
	"net/netip"
	"net/url"

	"github.com/shavac/mp1p/cfg"
	"github.com/shavac/mp1p/cmd"
	"github.com/shavac/mp1p/global"
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
	if err := cfg.ReadFromPath(cmd.CfgPath); err != nil {
		log.Errorln(err)
	}
	log.Setup()
	fmt.Println(global.GetConfig())
	log.Debugln(global.GetConfig().LogConfig.Encoding)
	for {
		//fmt.Println(cfg.Config())
		for portName, portCfg := range global.GetConfig().PortConfig {
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
				sCfg := global.GetConfig().ServiceConfig[sName]
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
