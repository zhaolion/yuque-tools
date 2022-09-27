/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"yuque-tools/cmd"
)

var (
	BuildStamp string
	Version    string
)

func init() {
	// 由于 ldflags -X 编译器似乎不能指定 main.go 以外的文件编译期添加参数，暂时在这里传递
	cmd.BuildStamp = BuildStamp
	cmd.Version = Version
}

func main() {
	cmd.Execute()
}
