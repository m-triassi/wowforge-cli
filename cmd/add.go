/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"
	"github.com/m-triassi/wowforge-cli/internal/search"
	"github.com/m-triassi/wowforge-cli/pkg/curseforge"
	"github.com/spf13/viper"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <addon id>",
	Short: "Add an addon to your list of tracked addons. ",
	Long: `Add an addon to your list of used addons. This wll immediately download and track
the addon of your choosing. The addon ID can be found under the "Project ID" heading in the "About Project" section of an 
addon's description

Be sure to only pass 1 addon id per addition. If you pass the id of an already tracked addon this command will simply 
update that addon in isolation.`,
	Run: func(cmd *cobra.Command, args []string) {
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
		if !search.Contains(list, modId) {
			list = append(list, modId)
			viper.Set("addons", list)
			viper.WriteConfig()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
