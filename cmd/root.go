/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/internal/flags"
	"github.com/alicemarple/lazydot/pkg/util/setup"
	"github.com/spf13/cobra"
)

var (
	PackageName string
	rootCmd     = &cobra.Command{
		Use:   "lazydot",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			setup.Setup(constants.ConfigFile)
			flags.FLagSetup(cmd, PackageName)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&PackageName, "sync", "S", "zsh", "Sync")
	rootCmd.Flags().StringVarP(&PackageName, "remove", "R", "zsh", "Remove")
	rootCmd.Flags().StringVarP(&PackageName, "search", "s", "all", "Search")
	rootCmd.Flags().BoolP("query", "Q", false, "Query")
	rootCmd.Flags().BoolP("update", "y", false, "Update")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
