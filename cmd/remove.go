/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/m-triassi/wowforge-cli/internal/files"
	"github.com/m-triassi/wowforge-cli/internal/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <addon id>",
	Short: "Remove the specified addon from game install",
	Long: `Removes the passed addon id from the tracked list of addons, and attempts to delete
the associated files for that addon.`,
	Run: func(cmd *cobra.Command, args []string) {
		modId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Errorf("passed mod ID is not strictly an integer: %w", err))
		}

		addons := search.LoadAddons()
		folders := addons[strconv.Itoa(modId)]
		delete(addons, strconv.Itoa(modId))

		installPath := viper.GetString("install")
		for _, folder := range folders {
			fmt.Printf("Deleting: %s... ", folder)
			del, err := filepath.Glob(filepath.Join(installPath, folder))
			if err != nil {
				panic(fmt.Errorf("could not read filesystem at path (%s): %w", installPath, err))
			}
			filesystem.DeleteAll(del)
			fmt.Printf("[DELETED]\n")
		}

		viper.Set("addons", addons)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
