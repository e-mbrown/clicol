package main

import (
	"fmt"
	"os"

	"github.com/clicol/base"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0][0] != '-' {
		//display help text
		fmt.Fprintln(os.Stderr, "Incorrect Args provided, here's help")
		os.Exit(1)
	}

	cfg, err := base.NewCliConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		os.Exit(1)
	}

	base.EntryPoint(args, cfg)
	fmt.Println("On new line")

}
