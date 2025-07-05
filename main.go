package main

import (
	"os"

	"github.com/ashish0kumar/cellscape/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
