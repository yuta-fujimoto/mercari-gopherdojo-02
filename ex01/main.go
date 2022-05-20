package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	}

	if err := run(args[0], 2); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
