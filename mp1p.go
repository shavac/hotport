package main

import (
	"net"

	"github.com/shavac/mp1p/cfg"
	"github.com/shavac/mp1p/cmd"
	"github.com/shavac/mp1p/log"
)

var (
	cfgChgEvt = make(chan bool)
)

func init() {
	cmd.Execute()
	cfg.OnChange(func() {
		cfgChgEvt <- true
	})
}

var allPorts = make(map[string]port)

type port struct {
	net.Listener
}

func main() {
	for {
		//fmt.Println(cfg.Config())
		for pName, pCfg := range cfg.Config().Port {
			l, err := net.Listen("tcp", pCfg.ListenAddr)
			if err != nil {
				log.S().Error(err)
			}
			allPorts[pName] = port{l}
			defer func(pName string) {
				allPorts[pName].Listener.Close()
				delete(allPorts, pName)
			}(pName)
		}
		<-cfgChgEvt
	}
}
