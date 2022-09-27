package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// version cmd
	Chain(rootCmd, "version", "v", version)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "Some tools to help me sync documents from yuque",
	Long:  `目前是一些简单的语雀文档 api 小工具`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
