package main

import (
	"os"

	"github.com/davdjl/testbeat/cmd"

	_ "github.com/davdjl/testbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
