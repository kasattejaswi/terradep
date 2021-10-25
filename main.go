package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kasattejaswi/terradep/deps"
)

var runFlag bool

func main() {
	flag.Parse()
	d, err := deps.ReadDcf("dependencyConfigs.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d.Dependencies)
}

func init() {
	flag.BoolVar(&runFlag, "run", false, "Start resolving dependencies. It must run in the root module")
}
