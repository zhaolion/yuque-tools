/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"yuque-tools/cmd"
)

func init() {
	cmd.BuildStamp = BuildStamp
	cmd.Version = Version
}

func main() {
	cmd.Execute()
}

var (
	BuildStamp string
	Version    string
)
