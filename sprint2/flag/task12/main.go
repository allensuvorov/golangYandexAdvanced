package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	// допишите код

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Version of %s\n%v\nUsage :\n", os.Args[0], version)
		flag.PrintDefaults()
	}

	flag.Parse()
}
