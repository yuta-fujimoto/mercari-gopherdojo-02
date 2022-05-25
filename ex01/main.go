package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	routineCnt = 2
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	}

	if err := Run(args[0], routineCnt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
