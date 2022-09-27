package cmd

import (
	"context"
	"fmt"
)

func user(_ context.Context, args []string) error {
	if len(args) == 0 {
		resp, err := client.UserCurrent()
		if err != nil {
			return err
		}

		fmt.Println(resp.String())
	}

	return nil
}
