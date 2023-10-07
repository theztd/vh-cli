/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"ztd/vh-cli/vashosting/dns"

	"github.com/spf13/cobra"
)

var zoneCmd = &cobra.Command{
	Use:     "zone",
	Aliases: []string{"z"},
	Short:   "Sprava DNS zon",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

var dnsListZones = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Vypis domen",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//outFmt, _ := cmd.Flags().GetString("out")
		zone, _ := cmd.Flags().GetString("zone")

		for id, r := range dns.ListZones(zone) {
			//fmt.Printf("%s - %+v\n", id, r)
			fmt.Printf("%s: %s\n", id, PrettyPrint(r))
		}
	},
}

func init() {

	zoneCmd.AddCommand(dnsListZones)
	rootCmd.AddCommand(zoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	zoneCmd.PersistentFlags().StringP("zone", "z", "", "DNS zone name")
	zoneCmd.PersistentFlags().StringP("out", "o", "", "Output format")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
