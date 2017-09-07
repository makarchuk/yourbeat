package main

import (
	"os"

	"github.com/makarchuk/yourbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
