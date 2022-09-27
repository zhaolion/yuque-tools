package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"yuque-tools/api"
)

var (
	authToken = ""
	client    *api.Client
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&authToken, "token", "t", "", "auth token to use")
	if v := os.Getenv("YUQUE_TOKEN"); v != "" && authToken == "" {
		authToken = v
	}

	// version cmd
	Chain(rootCmd, "version", "show binary version", version)
	// hello cmd, test the token
	Chain(rootCmd, "hello", "test token", hello)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "Some tools to help me sync documents from yuque",
	Long:  `目前是一些简单的语雀文档 api 小工具`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = api.NewClient(authToken)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
