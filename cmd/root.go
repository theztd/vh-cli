/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"ztd/vh-cli/vashosting"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vh-cli",
	Version: vashosting.VERSION,
	Short:   "Vas-Hosting CLI client",
	Long: fmt.Sprintf(`                           
             ▅▅        
           ▅▅▅▅▅▅          Vas Hosting
            ▅▅▅▅           (%s)
             ▅▅

🌎 Command Line Interface pro komunikaci s Vas Hosting API
	`, vashosting.VERSION),
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
