package cmd

import (
	"context"
	"fmt"
)

func user(_ context.Context, args []string) error {
	resp, err := client.User(args...)
	if err != nil {
		return err
	}

	fmt.Println(resp.String())

	return nil
}
