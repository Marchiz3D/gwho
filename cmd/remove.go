/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [alias]",
	Short: "Remove a saved Git profile",
	Long: `Remove an existing Git profile from your configuration by its alias.
This action is permanent and cannot be undone.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: You must provide an alias (e.g., 'gwho remove work')")
			return
		}
		alias := args[0]

		config, err := LoadConfig()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		if _, ok := config.Profiles[alias]; !ok {
			fmt.Println("Error: Alias '" + alias + "' does not exist")
			return
		}

		delete(config.Profiles, alias)

		if err := SaveConfig(config); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

		fmt.Printf("Profile '%s' removed successfully!\n", alias)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
