/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [alias]",
	Short: "Edit an existing Git profile",
	Long: `Edit an existing Git profile interactively.

You must provide the alias of the profile you want to edit.
The CLI will display the current name and email, then prompt you for new values.
To keep the existing value for a field, simply press Enter without typing anything.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("Error: You must provide an alias (e.g., 'gwho edit work')\n")
			return
		}
		alias := args[0]

		config, err := LoadConfig()
		if err != nil {
			color.Red("Error loading config: %v\n", err)
			return
		}

		if _, ok := config.Profiles[alias]; !ok {
			color.Red("Error: Alias '%s' does not exist\n", alias)
			return
		}

		render := bufio.NewReader(os.Stdin)
		color.Yellow("Current Name: %s, Email: %s\n", config.Profiles[alias].Name, config.Profiles[alias].Email)
		color.Cyan("Enter Name: ")
		name, _ := render.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = config.Profiles[alias].Name
		}

		color.Cyan("Enter Email: ")
		email, _ := render.ReadString('\n')
		email = strings.TrimSpace(email)
		if email == "" {
			email = config.Profiles[alias].Email
		}

		if name == "" || email == "" {
			color.Red("Error: Name and email cannot be empty\n")
			return
		}

		if !strings.Contains(email, "@") {
			color.Red("Error: Invalid email\n")
			return
		}

		config.Profiles[alias] = Profile{
			Name:  name,
			Email: email,
		}

		if err := SaveConfig(config); err != nil {
			color.Red("Error saving config: %v\n", err)
			return
		}

		color.Green("> Profile '%s' updated successfully!\n", alias)
		cyanBold := color.New(color.FgCyan, color.Bold)
		cyanBold.Printf("Name: %s\nEmail: %s\n", name, email)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
