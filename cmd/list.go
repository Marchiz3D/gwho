/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sort"

	"github.com/Marchiz3D/gwho/service"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved Git profiles",
	Long: `Display a list of all Git profiles currently saved in your configuration.
Each profile will show its alias, name, and email address.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := service.GetAllProfiles()
		if err != nil {
			color.Red("Error: %v\n", err)
			return
		}

		color.Cyan("\n📋 List of Git Profiles:")
		color.Cyan("========================")

		aliases := make([]string, 0, len(res))
		for alias := range res {
			aliases = append(aliases, alias)
		}

		sort.Strings(aliases)

		for i, alias := range aliases {
			profile := res[alias]
			greenBold := color.New(color.FgGreen, color.Bold)
			greenBold.Printf("[%d] %s\n", i+1, alias)
			color.Magenta("Name: %s\nEmail: %s\n", profile.Name, profile.Email)
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
