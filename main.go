package main

import "github.com/KvalitetsIT/kih-telecare-exporter/cmd"

var (
	Build   string
	Version string
)

func main() {
	cmd.Build = Build
	cmd.Version = Version
	cmd.Execute()
}
