package main

import (
	"os"

	"github.com/stratumfarm/derocli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
