package main

import (
	"os"

	"github.com/jackcloudman/testbeat/cmd"

	_ "github.com/jackcloudman/testbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
