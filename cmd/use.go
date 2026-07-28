/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use [alias]",
	Short: "Switch to a saved Git profile",
	Long: `Apply a saved Git profile to the current local repository.
This command runs 'git config' under the hood to set the user.name and user.email
for the current project based on the alias provided.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		config, _ := LoadConfig()
		aliases := make([]string, 0, len(config.Profiles))
		for alias := range config.Profiles {
			aliases = append(aliases, alias)
		}
		return aliases, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		config, err := LoadConfig()
		if err != nil {
			color.Red("Error loading config: %v\n", err)
			return
		}

		var selectedAlias string
		if len(args) > 0 {
			selectedAlias = args[0]
		} else {
			aliases := make([]string, 0, len(config.Profiles))
			for alias := range config.Profiles {
				aliases = append(aliases, alias)
			}

			sort.Strings(aliases)
			for i, alias := range aliases {
				idxStr := color.HiYellowString("[%d]", i+1)
				aliasStr := color.HiGreenString(alias)
				fmt.Printf("%s %s\n", idxStr, aliasStr)
			}

			var choice int
			color.Cyan("Select a profile: ")
			_, err := fmt.Scanln(&choice)

			if err != nil {
				color.Red("Error: Invalid choice\n")
				return
			}

			if choice < 1 || choice > len(aliases) {
				color.Red("Error: Invalid choice\n")
				return
			}

			selectedAlias = aliases[choice-1]
		}

		profile, ok := config.Profiles[selectedAlias]
		if !ok {
			color.Red("Error: Profile '%s' not found\n", selectedAlias)
			return
		}

		var err1, err2 error
		isGlobal, _ := cmd.Flags().GetBool("global")

		if isGlobal {
			err1 = exec.Command("git", "config", "--global", "user.name", profile.Name).Run()
			err2 = exec.Command("git", "config", "--global", "user.email", profile.Email).Run()

			color.Green("> Switched to global profile '%s'\n", selectedAlias)

		} else {
			err1 = exec.Command("git", "config", "user.name", profile.Name).Run()
			err2 = exec.Command("git", "config", "user.email", profile.Email).Run()

			color.Green("> Switched to local profile '%s'\n", selectedAlias)
		}

		if err1 != nil || err2 != nil {
			color.Red("Error applying profile. Make sure that you're in a git repository.\n")
			return
		}

		cyanBold := color.New(color.FgCyan, color.Bold)
		cyanBold.Printf("Name: %s\nEmail: %s\n", profile.Name, profile.Email)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.
	useCmd.Flags().BoolP("global", "g", false, "Apply the git config globally")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
