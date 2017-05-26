package main

import (
	"./log_emitter"
	"flag"
	"fmt"
)

const (
	DEFAULT_CONF_FILE = "conf.yml"
)

var (
	Version string
	Build   string
)

func main() {
	fmt.Println("Version: ", Version)
	fmt.Println("Build: ", Build)

	confPtr := flag.String("conf", DEFAULT_CONF_FILE, "path to conf file")
	flag.Parse()

	var le le.LogEmitter
	le.LoadConf(*confPtr)
	le.Execute()
}
