/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"
	"github.com/m-triassi/wowforge-cli/pkg/curseforge"
	"github.com/spf13/viper"
	"strconv"

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
		modId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Errorf("Passed mod ID is not strictly an integer: %w", err))
		}

		files, err := curseforge.GetFiles(modId)
		if err != nil {
			panic(fmt.Errorf("Failed to fetch files from CurseForge, mod id may not exist: %w", err))
		}

		file, err := curseforge.DownloadFile(modId, curseforge.NegotiateFile(files))
		if err != nil {
			panic(fmt.Errorf("Could not download file: %w", err))
		}

		dest := viper.GetString("install")

		err = curseforge.InstallAddon(file, dest)
		if err != nil {
			panic(fmt.Errorf("Failed to install addon in target destination"))
		}

		list := viper.GetIntSlice("addons")
		if !contains(list, modId) {
			list = append(list, modId)
			viper.Set("addons", list)
			viper.WriteConfig()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func contains(haystack []int, needle int) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
