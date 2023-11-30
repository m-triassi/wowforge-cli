/*
Copyright © 2023 Massimo Triassi <contact@triassi.ca>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wowforge-cli",
	Short: "A light installer for fetching your latest addons for WoW",
	Long: `wowforge-cli is a light-weight command line utility that can fetch the latest versions
of your chosen addons. It maintains a list of tracked addons for you, and allows you to update them
all together.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wowforge-cli.json)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Fprintln(os.Stderr, "Using config file:", cfgFile)
	} else {
		//Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".wowforge-cli" (without extension).
		viper.SetConfigName(".wowforge-cli")
		viper.SetConfigType("json")
		viper.AddConfigPath(home)

	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error Reading config file:", err)
	}
}
