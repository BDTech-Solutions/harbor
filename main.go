package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/BDTech-Solutions/harbor/cmd"
)

//go:embed VERSION
var versionFile string

func main() {
	version := strings.TrimSpace(versionFile)
	root := cmd.NewRootCommand(version)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "❌ "+err.Error())
		os.Exit(1)
	}
}
