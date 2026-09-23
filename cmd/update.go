/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"
	"github.com/m-triassi/wowforge-cli/internal/search"
	"github.com/m-triassi/wowforge-cli/pkg/curseforge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all tracked addons",
	Long: `Fetches the latest version of all addons currently being tracked by wowforge-cli.
Simply run the command and all files will be downloaded and unpacked.
`,
	Run: func(cmd *cobra.Command, args []string) {
		addons := search.LoadAddons()

		for key := range addons {
			id, _ := strconv.Atoi(key)
			files, err := curseforge.GetFiles(id)
			if err != nil {
				panic(fmt.Errorf("Failed to fetch files from CurseForge, mod id (%d) may not exist: %w", id, err))
			}

			negotiated := curseforge.NegotiateFile(files, viper.GetString("flavor"))
			if negotiated.Id == 0 {
				fmt.Printf("\nNo compatible file found for addon %d, skipping.\n", id)
				continue
			}
			fmt.Printf("\nDownloading: %s... ", negotiated.Filename)
			file, err := curseforge.DownloadFile(id, negotiated)
			if err != nil {
				panic(fmt.Errorf("Could not download file: %w", err))
			}

			fmt.Printf("Unpacking... ")
			dest := viper.GetString("install")
			folders, err := curseforge.InstallAddon(file, dest)
			if err != nil {
				panic(fmt.Errorf("Failed to install addon in target destination"))
			}

			addons[key] = folders
			viper.Set("addons", addons)
			viper.WriteConfig()
			fmt.Printf("[INSTALLED]\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
