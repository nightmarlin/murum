package main

import (
	"fmt"
	"os"

	"github.com/nightmarlin/murum/cmd"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
