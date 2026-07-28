/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
			fmt.Println("Error: You must provide an alias (e.g., 'gwho add work')")
			return
		}
		alias := args[0]

		render := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Name: ")
		name, _ := render.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Enter Email: ")
		email, _ := render.ReadString('\n')
		email = strings.TrimSpace(email)

		if name == "" || email == "" {
			fmt.Println("Error: Name and email cannot be empty")
			return
		}

		if !strings.Contains(email, "@") {
			fmt.Println("Error: Invalid email")
			return
		}

		config, err := LoadConfig()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		if _, ok := config.Profiles[alias]; ok {
			fmt.Println("Error: Alias '" + alias + "' already exists")
			return
		}

		config.Profiles[alias] = Profile{
			Name:  name,
			Email: email,
		}

		if err := SaveConfig(config); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

		fmt.Printf("> Profile '%s' added successfully!\n", alias)
		fmt.Printf("Name: %s\nEmail: %s\n", name, email)
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
