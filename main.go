package main

import (
	"os"

  "github.com/DrLivsey00/transaction-parcer-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
