/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ztd/vh-cli/cmd/dns"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var dnsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List DNS records in specified zone",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		zone, _ := cmd.Flags().GetString("zone")
		dns.List(zone)
	},
}

func init() {
	dnsCmd.AddCommand(dnsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
