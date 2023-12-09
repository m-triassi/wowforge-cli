/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration values for the operation of the tool.",
	Long: `Set required values for the operation of the tool. This include setting things like installation path
of World of Warcraft, or your CurseForge API key (where required).

Note, when specifying things like install directories, if the path would contain a space, be sure to either encapsulate
the path in quotations or escape all spaces, like so:

  wowforge-cli set --install "/path to/my/install location"
  wowforge-cli set --install /path\ to/my/install\ location
`,
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := os.Stat(viper.ConfigFileUsed()); errors.Is(err, os.ErrNotExist) {
			os.Create(viper.ConfigFileUsed())
		}

		writeErr := viper.WriteConfig()
		if writeErr != nil {
			fmt.Println("Error writing settings:", writeErr)
		}

		settings, readErr := json.MarshalIndent(viper.AllSettings(), "", "  ")
		if readErr != nil {
			fmt.Println("Error reading settings:", readErr)
		}

		fmt.Println(string(settings))
	},
	Args: cobra.NoArgs,
}

var Install string
var ApiKey string
var DownloadDir string

func init() {
	rootCmd.AddCommand(setCmd)

	apiUsage := "API key from your CurseForge Studio account. This is required for some endpoints to function correctly"
	installUsage := "Path to your target installation of World of Warcraft"
	downloadUsage := "Path for downloaded addons to be stored (default: /tmp/)"

	originalInstall, installErr := setCmd.Flags().GetString("install")
	if installErr == nil {
		fmt.Println("Could not get value for \"install\" configuration value.", installErr)
	}

	originalApi, apiErr := setCmd.Flags().GetString("api-key")
	if apiErr == nil {
		fmt.Println("Could not get value for \"api-key\" configuration value.", apiErr)
	}

	originalDownload, downloadErr := setCmd.Flags().GetString("download-dir")

	if downloadErr == nil {
		fmt.Println("Could not get value for \"download-dir\" configuration value.", downloadErr)
	}

	setCmd.Flags().StringVar(&ApiKey, "api-key", originalApi, apiUsage)
	setCmd.Flags().StringVar(&DownloadDir, "download-dir", originalDownload, downloadUsage)
	setCmd.Flags().StringVar(&Install, "install", originalInstall, installUsage)

	if apiErr := viper.BindPFlag("api-key", setCmd.Flags().Lookup("api-key")); apiErr != nil {
		fmt.Println(apiErr)
	}

	if downloadErr := viper.BindPFlag("download-dir", setCmd.Flags().Lookup("download-dir")); downloadErr != nil {
		fmt.Println(downloadErr)
	}

	if installErr := viper.BindPFlag("install", setCmd.Flags().Lookup("install")); installErr != nil {
		fmt.Println(installErr)
	}

}
