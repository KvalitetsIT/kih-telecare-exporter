package main

import "bitbucket.org/opentelehealth/exporter/cmd"

var (
	Build   string
	Version string
)

func main() {
	cmd.Build = Build
	cmd.Version = Version
	cmd.Execute()
}
