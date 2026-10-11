package main

import (
	"cut/internal/cli"
	"os"
)

func main() {
	os.Exit(cli.Run(os.Args, os.Stdout, os.Stderr))
}
