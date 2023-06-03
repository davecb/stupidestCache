package main

// stupid -- a program to exercise the simplest/stupidest possible cache.

import (
	"errors"
	"flag"
	"fmt"
	"github.com/davecb/stupidestCache/src/fromFile"
	"github.com/davecb/stupidestCache/src/fromHttp"
	"log"
	"os"
)

var eShort = errors.New("csv record was too short")
var eBlank = errors.New("csv record was blank")

// main -- get options and commands
func main() {
	var daemonic = flag.Bool("daemonic", false, "bring up a web server on port 8080")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime) // show fromFile:line in logs

	flag.Parse()
	if *daemonic {
		// run it as a deamon
		fromHttp.Run()
		return
	}
	if flag.NArg() < 1 {
		fmt.Fprint(os.Stderr, "You must supply a load.csv file\n") //nolint
		usage()
		os.Exit(1)
	}
	filename := flag.Arg(0)
	fromFile.Run(filename)
}

func usage() {
	//nolint
	fmt.Fprint(os.Stderr, "Usage: stupid fromFile.csv") //nolint
	flag.PrintDefaults()
}
