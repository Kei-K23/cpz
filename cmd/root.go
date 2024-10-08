/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cpz",
	Version: "0.1.0",
	Short:   "A fast, modern and batteries include 'cp/mv' alternative or replacement terminal command written in Go",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func helpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Help about any command",
	}
}

func init() {
	completion := completionCommand()
	// mark completion command hidden
	completion.Hidden = true

	rootCmd.AddCommand(completion)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
