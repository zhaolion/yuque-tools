package cmd

import (
	"context"
	"fmt"
)

func hello(_ context.Context, _ []string) error {
	resp, err := client.Hello()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s", *resp.Data.Message))
	return nil
}
