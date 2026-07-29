/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/Marchiz3D/gwho/service"
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

		profiles, _ := service.GetAllProfiles()
		aliases := make([]string, 0, len(profiles))
		for alias := range profiles {
			aliases = append(aliases, alias)
		}
		return aliases, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		isGlobal, _ := cmd.Flags().GetBool("global")

		var selectedAlias string
		if len(args) > 0 {
			selectedAlias = args[0]
		} else {
			profiles, err := service.GetAllProfiles()
			if err != nil {
				color.Red("Error: %v\n", err)
				return
			}

			if len(profiles) == 0 {
				color.Yellow("No Git profiles found. Use 'gwho add <alias>' to create one.\n")
				return
			}

			aliases := make([]string, 0, len(profiles))
			for alias := range profiles {
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
			_, err = fmt.Scanln(&choice)

			if err != nil || choice < 1 || choice > len(aliases) {
				color.Red("Error: Invalid choice\n")
				return
			}
			selectedAlias = aliases[choice-1]
		}

		profile, err := service.UseProfile(selectedAlias, isGlobal)
		if err != nil {
			color.Red("%v\n", err)
			return
		}

		if isGlobal {
			color.Green("> Switched to global profile '%s'\n", selectedAlias)
		} else {
			color.Green("> Switched to local profile '%s'\n", selectedAlias)
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
