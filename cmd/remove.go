/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Marchiz3D/gwho/service"
	"github.com/fatih/color"
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
			color.Red("Error: You must provide an alias (e.g., 'gwho remove work')\n")
			return
		}
		alias := args[0]

		if err := service.DeleteProfile(alias); err != nil {
			color.Red("Error: %v\n", err)
			return
		}

		color.Green("> Profile '%s' removed successfully!\n", alias)
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
