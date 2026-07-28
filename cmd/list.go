/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

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
		config, err := LoadConfig()
		if err != nil {
			color.Red("Error loading config: %v\n", err)
			return
		}

		if len(config.Profiles) == 0 {
			color.Yellow("No Git profiles found. Use 'gwho add <alias>' to create one.")
			return
		}

		color.Cyan("\n📋 List of Git Profiles:")
		color.Cyan("========================")

		aliases := make([]string, 0, len(config.Profiles))
		for alias := range config.Profiles {
			aliases = append(aliases, alias)
		}

		sort.Strings(aliases)

		for i, alias := range aliases {
			profile := config.Profiles[alias]

			idxStr := color.HiYellowString("[%d]", i+1)
			aliasStr := color.HiGreenString(alias)
			emailStr := color.HiBlackString("<%s>", profile.Email)

			fmt.Printf("%s %s - %s %s\n", idxStr, aliasStr, profile.Name, emailStr)
		}
		fmt.Println()
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
