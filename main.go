package main

import (
	"flag"
)

var runFlag bool

func main() {
	flag.Parse()
	// if !runFlag {
	// 	fmt.Println(`
	// 	Usage terradep command [args]
	// 	run					Start resolv
	// 	`)

	// }
}

func init() {
	flag.BoolVar(&runFlag, "run", false, "Start resolving dependencies. It must run in the root module")
}
