package base

import (
	"flag"
	"fmt"

	"github.com/clicol/proc"
)

type FlagNameSet struct {
	fs      *flag.FlagSet
	NameMap map[string]string
}

// PopulateDefFlagSet makes a new flag set and also has a map to capture data to process.
func PopulateDefFlagSet() (FlagNameSet, error) {
	fns := FlagNameSet{
		fs:      flag.NewFlagSet("Default Flags", flag.ContinueOnError),
		NameMap: map[string]string{},
	}

	fns.NameMap["-prs"] = "="
	fns.fs.BoolFunc("prs", "`Process` will look at listed images and return info", proc.Process)

	fns.NameMap["-test"] = "="
	fns.fs.BoolFunc("test", "We be testing things with this command", func(s string) error {
		fmt.Println("Test and test arg:", s)
		return nil
	})

	return fns, nil
}
