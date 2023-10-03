/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"ztd/vh-cli/vashosting/dns"

	"github.com/spf13/cobra"
)

var dnsList = &cobra.Command{
	Use:   "list",
	Short: "List zaznamu v domene",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		outFmt, _ := cmd.Flags().GetString("out")
		zone, _ := cmd.Flags().GetString("zone")
		fmt.Println("dnsList called. Output format will be", outFmt)

		for id, r := range dns.List(zone) {
			//fmt.Printf("%s - %+v\n", id, r)
			fmt.Printf("%s: %s\n", id, PrettyPrint(r))
		}
	},
}

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Sprava DNS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

func init() {
	dnsCmd.AddCommand(dnsList)
	rootCmd.AddCommand(dnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	dnsCmd.PersistentFlags().StringP("zone", "z", "", "DNS zone name")
	dnsCmd.PersistentFlags().StringP("out", "o", "", "Output format")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
