package main

import (
	"os"

	"github.com/gofrontier-com/go-utils/output"
	"github.com/gofrontier-com/pony-express/pkg/cmd/pony"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	command := pony.NewRootCmd(version, commit, date)
	if err := command.Execute(); err != nil {
		output.PrintlnError(err)
		os.Exit(1)
	}
}
