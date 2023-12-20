/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/m-triassi/wowforge-cli/internal/files"
	"github.com/m-triassi/wowforge-cli/internal/search"
	"github.com/m-triassi/wowforge-cli/pkg/curseforge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"regexp"
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

		addons := viper.GetIntSlice("addons")
		remove := search.Find(addons, modId)
		addons = append(addons[:remove], addons[remove+1:]...)

		fileSet, _ := curseforge.GetFiles(modId)
		filename := curseforge.NegotiateFile(fileSet).Filename
		re := regexp.MustCompile("[a-zA-Z]*")
		res := string(re.Find([]byte(filename)))

		del, err := filepath.Glob(viper.GetString("install") + res + "*")

		filesystem.DeleteAll(del)

		viper.Set("addons", addons)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
