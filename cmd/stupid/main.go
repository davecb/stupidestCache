package main

// stupid -- a program to exercise the simplest/stupidest possible cache.

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/davecb/stupidestCache/src/fromFile"
	"github.com/davecb/stupidestCache/src/fromHttp"
)

var eShort = errors.New("csv record was too short")
var eBlank = errors.New("csv record was blank")

// main -- get options and commands
func main() {
	var daemonic = flag.Bool("daemonic", false, "bring up a web server on port 8080")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime) // show fromFile:line in logs

	flag.Parse()

	if flag.NArg() < 1 && !*daemonic {
		fmt.Fprint(os.Stderr, "You must supply a load.csv file\n") //nolint
		usage()
		os.Exit(1)
	}
	// In all cases, interpret all the files on the command line
	for i := range flag.Args() {
		fromFile.Run(flag.Arg(i))
	}

	if *daemonic {
		// run it as a daemon after any files
		fromHttp.Run()
		return
	}
}

func usage() {
	//nolint
	fmt.Fprint(os.Stderr, "Usage: stupid fromFile.csv") //nolint
	flag.PrintDefaults()
}
