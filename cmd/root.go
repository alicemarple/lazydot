/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	packageName string
	BaseURL     = "https://github.com/"
	// BaseURL = "https://github.com/alicemarple/dotfiles/releases/download"
	rootCmd = &cobra.Command{
		Use:   "lazydot",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runFun(cmd)
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
	rootCmd.Flags().StringVarP(&packageName, "sync", "S", "zsh", "Sync")
	rootCmd.Flags().StringVarP(&packageName, "remove", "R", "zsh", "Remove")
	rootCmd.Flags().BoolP("query", "Q", false, "Query")
	rootCmd.Flags().BoolP("update", "Y", false, "Update")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runFun(cmd *cobra.Command) {
	BaseURL = fmt.Sprintf("%s/", BaseURL)
	isSync := cmd.Flags().Changed("sync")
	isRemove := cmd.Flags().Changed("remove")
	isQuery := cmd.Flags().Changed("query")
	isUpdate := cmd.Flags().Changed("update")

	if isSync {
		sync("/mnt/e/projects/golang/lazydot/sync/sync.yml", packageName)
	} else if isRemove {
		remove(packageName)
	} else if isQuery {
		query()
	} else if isUpdate {
		update()
	} else {
		fmt.Println("Give the proper flag")
	}
}
