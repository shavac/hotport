package main

import (
	"net"

	"github.com/prometheus/common/log"
	"github.com/shavac/mp1p/cfg"
	"github.com/shavac/mp1p/cmd"
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
				log.Errorln(err)
			}
			allPorts[pName] = port{l}
			defer func() {
				allPorts[pName].Listener.Close()
				delete(allPorts, pName)
			}()
		}
		<-cfgChgEvt
	}
}
