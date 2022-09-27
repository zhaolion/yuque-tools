package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

type CommandHandle func(ctx context.Context, args []string) error

type CommandOption func(*cobra.Command)

// Chain 绑定子命令
func Chain(mainCmd *cobra.Command, use, short string, handle CommandHandle, options ...CommandOption) {
	command := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			err := handle(ctx, args)
			// 反馈报错信号
			if err != nil {
				log.Fatalf("err: %+v", err)
			}
		},
	}

	for _, opt := range options {
		opt(command)
	}

	mainCmd.AddCommand(command)
}
