package main

import (
	_ "embed"
	"strings"

	"github.com/BDTech-Solutions/harbor/cmd"
)

//go:embed VERSION
var versionFile string

func main() {
	version := strings.TrimSpace(versionFile)
	cmd.Execute(version)
}
