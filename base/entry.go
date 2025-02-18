package base

import (
	"fmt"
	"os"
)

func EntryPoint(args []string, cfg CliConf) {
	var currFlag string

	defFlags, err := PopulateDefFlagSet()
	if err != nil {
		fmt.Fprintf(cfg.ErrStream, "Error creating flagset: %v\n", err)
		os.Exit(1)
	}
	defFlags.fs.SetOutput(cfg.OutStream)

	for _, word := range args {
		if word[0] == '-' {
			currFlag = word
			if _, ok := defFlags.NameMap[currFlag]; !ok {
				// Flag doesnt exist
				fmt.Fprintf(cfg.ErrStream, "Error flag doesnt exists: %v\n", word)
				os.Exit(1)
			}
		} else {
			defFlags.NameMap[currFlag] += word + " "
		}
	}

	// TODO: Find another way to parse flags, maybe an arr with names encounter.)
	for k, v := range defFlags.NameMap {
		//TODO: capture error if continue on error
		defFlags.fs.Parse([]string{k + v})
	}

}
