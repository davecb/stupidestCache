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
	var daemonic = flag.Bool("daemonic", false, "bring up a web server on port 80")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime) // show fromFile:line in logs

	flag.Parse()
	if *daemonic {
		// run it as a deamon
		fromHttp.Run()
	} else {
		if flag.NArg() < 1 {
			fmt.Fprint(os.Stderr, "You must supply a load.csv fromFile\n") //nolint
			usage()
		}

		filename := flag.Arg(0)
		fromFile.Run(filename)

	}

}

func usage() {
	//nolint
	fmt.Fprint(os.Stderr, "Usage: stupid fromFile.csv") //nolint
	flag.PrintDefaults()
	os.Exit(1)
}
