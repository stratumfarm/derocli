package main

import (
	"log"

	"github.com/stratumfarm/derocli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
