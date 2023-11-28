/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Search for, and add an addon to your tracked addons",
	Long: `Add an addon to your list of used addons. This wll immediately download and track
the addon of your choosing. Usage:

wowforge-cli add [addon name]

Be sure to only pass 1 addon name per addition. You will be prompted to choose which addon from the result set you'd like`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

var Addon string

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&Addon, "id", "i", "", "Specify the addon by it CurseForge {modId}. Avoids interactive prompts.")
}
