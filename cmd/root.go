/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gwho",
	Short: "A simple Git profile switcher",
	Long: `gwho (Gw Who? | Indonesian Slang: Who Am I?) is a CLI tool to seamlessly manage and switch 
between multiple Git profiles (name and email) across different workspaces.

It helps prevent committing with the wrong identity by allowing you to easily
verify your current profile and switch to the correct one.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tip: Use 'gwho --help' to see all available commands.")
		fmt.Println("---")
		fmt.Println("[TODO: Display the currently active Git profile here]")
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gwho.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
