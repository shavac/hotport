package main

import (
	"fmt"
	"net"
	"net/netip"
	"net/url"

	"github.com/shavac/hotport/cfg"
	"github.com/shavac/hotport/cmd"
	"github.com/shavac/hotport/global"
	"github.com/shavac/hotport/log"
	"github.com/shavac/hotport/port"
	"github.com/shavac/hotport/proto"
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
	log.Debugln(global.GetConfig().Log.Encoding)
	for {
		//fmt.Println(cfg.Config())
		for portName, portCfg := range global.GetConfig().Port {
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
				sCfg := global.GetConfig().Service[sName]
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
