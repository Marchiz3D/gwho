/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the active Git profile",
	Long: `Identify and display the Git profile currently active in this repository.
It reads the local Git configuration and matches it against your saved profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := LoadConfig()
		if err != nil {
			color.Red("Error loading config: %v\n", err)
			return
		}

		currentName := getGitConfig("user.name")
		currentEmail := getGitConfig("user.email")

		for alias, profile := range config.Profiles {
			if profile.Name == currentName && profile.Email == currentEmail {
				color.Green("> You are currently using '%s' profile\n", alias)
				cyanBold := color.New(color.FgCyan, color.Bold)
				cyanBold.Printf("Name: %s\nEmail: %s\n", profile.Name, profile.Email)
				return
			}
		}

		color.Yellow("You are currently not using any profile.\nUse 'gwho add <alias>' to add a profile or 'gwho list' to list all available profiles.\n")
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
