/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/Marchiz3D/gwho/config"
	"github.com/Marchiz3D/gwho/service"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [alias]",
	Short: "Add a new Git profile",
	Long: `Add a new Git profile to your configuration interactively.

You must provide a unique alias for the profile (e.g., 'work', 'personal').
The CLI will then prompt you to enter the associated name and email address.
These profiles can later be applied to any repository using 'gwho use <alias>'.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("Error: You must provide an alias (e.g., 'gwho add work')\n")
			return
		}
		alias := args[0]

		render := bufio.NewReader(os.Stdin)
		color.Cyan("Enter Name: ")
		name, _ := render.ReadString('\n')
		name = strings.TrimSpace(name)

		color.Cyan("Enter Email: ")
		email, _ := render.ReadString('\n')
		email = strings.TrimSpace(email)

		if name == "" || email == "" {
			color.Red("Error: Name and email cannot be empty\n")
			return
		}

		if !strings.Contains(email, "@") {
			color.Red("Error: Invalid email\n")
			return
		}

		profile := config.Profile{
			Name:  name,
			Email: email,
		}

		if err := service.CreateProfile(alias, profile); err != nil {
			color.Red("Error: %v\n", err)
			return
		}

		color.Green("> Profile '%s' added successfully!\n", alias)
		cyanBold := color.New(color.FgCyan, color.Bold)
		cyanBold.Printf("Name: %s\nEmail: %s\n", name, email)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
