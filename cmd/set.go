/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println(cmd.Flags())
	},
	Args: cobra.NoArgs,
}

var Install string
var ApiKey string

func init() {
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVar(&ApiKey, "api-key", "", "API key from your CurseForge Studio account. This is required for some endpoints to function correctly")
	setCmd.Flags().StringVar(&Install, "install", "", "path to your target installation of World of Warcraft")

	if apiErr := viper.BindPFlag("api-key", setCmd.Flags().Lookup("api-key")); apiErr != nil {
		return
	}

	if installErr := viper.BindPFlag("install", setCmd.Flags().Lookup("install")); installErr != nil {
		return
	}
}
