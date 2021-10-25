package main

import (
	"flag"

	"github.com/kasattejaswi/terradep/deps"
)

var runFlag bool

func main() {
	flag.Parse()
	deps.Greet()
}

func init() {
	flag.BoolVar(&runFlag, "run", false, "Start resolving dependencies. It must run in the root module")
}
