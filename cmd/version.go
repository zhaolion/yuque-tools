package cmd

import (
	"context"
	"fmt"
)

var (
	BuildStamp string
	Version    string
)

func version(_ context.Context, _ []string) error {
	fmt.Println(fmt.Sprintf("version: %s", Version))
	fmt.Println(fmt.Sprintf("build stamp: %s", BuildStamp))
	return nil
}
