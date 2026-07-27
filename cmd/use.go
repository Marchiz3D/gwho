/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use [alias]",
	Short: "Switch to a saved Git profile",
	Long: `Apply a saved Git profile to the current local repository.
This command runs 'git config' under the hood to set the user.name and user.email
for the current project based on the alias provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: You must provide an alias (e.g., 'gwho use work')")
			return
		}
		alias := args[0]

		config, err := LoadConfig()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		profile, ok := config.Profiles[alias]
		if !ok {
			fmt.Printf("Error: Profile '%s' not found\n", alias)
			return
		}

		err1 := exec.Command("git", "config", "user.name", profile.Name).Run()
		err2 := exec.Command("git", "config", "user.email", profile.Email).Run()

		if err1 != nil || err2 != nil {
			fmt.Println("Error applying profile. Make sure that you're in a git repository.")
			return
		}

		fmt.Printf("Switched to profile '%s'\n", alias)
		fmt.Printf("Name: %s\nEmail: %s\n", profile.Name, profile.Email)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
