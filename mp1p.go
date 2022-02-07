package main

import (
	"flag"
	"log"

	"github.com/shavac/mp1p/cfg"
	"github.com/shavac/mp1p/cmd"
)

var (
	cfgPath *string
)

func init() {
	flag.Parse()
	cfgPath = flag.String("c", "/etc/mp1p/mp1p.toml", "config file")

}

func main() {
	cmd.Execute()
	c, err := cfg.ReadFromTomlFile("etc/mp1p/mp1p.toml")
	if err != nil {
		log.Fatalln(err)
	}
	_ = c
	//fmt.Printf("%+v", c)
}
