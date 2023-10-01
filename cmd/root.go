/*
Copyright © 2023 Marek Sirovy
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vh-cli",
	Short: "Vas-Hosting CLI client",
	Long: `        
            Vas Hosting
        ==================

                 ▅▅
               ▅▅▅▅▅▅
                ▅▅▅▅
                 ▅▅


🌎 Command Line Interface pro komunikaci s Vas Hosting API
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vh-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
