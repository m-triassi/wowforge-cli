/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"
	"github.com/m-triassi/wowforge-cli/pkg/curseforge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all tracked addons",
	Long: `Fetches the latest version of all addons currently being tracked by wowforge-cli. 
Simply run the command and all files will be downloaded and unpacked.
`,
	Run: func(cmd *cobra.Command, args []string) {
		addons := viper.GetIntSlice("addons")

		for _, id := range addons {
			files, err := curseforge.GetFiles(id)
			if err != nil {
				panic(fmt.Errorf("Failed to fetch files from CurseForge, mod id (%d) may not exist: %w", id, err))
			}

			negotiated := curseforge.NegotiateFile(files)
			fmt.Printf("\nDownloading: %s... ", negotiated.Filename)
			file, err := curseforge.DownloadFile(id, negotiated)
			if err != nil {
				panic(fmt.Errorf("Could not download file: %w", err))
			}

			fmt.Printf("Unpacking... ")
			dest := viper.GetString("install")
			err = curseforge.InstallAddon(file, dest)
			if err != nil {
				panic(fmt.Errorf("Failed to install addon in target destination"))
			}

			fmt.Printf("[INSTALLED]\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
