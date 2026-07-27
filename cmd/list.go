/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved Git profiles",
	Long: `Display a list of all Git profiles currently saved in your configuration.
Each profile will show its alias, name, and email address.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List of Git Profiles:")
		config, err := LoadConfig()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		for alias, profile := range config.Profiles {
			fmt.Printf("%s - %s <%s>\n", alias, profile.Name, profile.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
