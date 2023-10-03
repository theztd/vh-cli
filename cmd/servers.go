/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"ztd/vh-cli/vashosting/servers"

	"github.com/spf13/cobra"
)

// serversCmd represents the servers command
var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "Manage servers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

var serversList = &cobra.Command{
	Use:   "list",
	Short: "List all servers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List all servers")
		name, _ := cmd.Flags().GetString("name")
		// clr := ""
		for id, r := range servers.List(name) {
			//fmt.Printf("%s - %+v\n", id, r)
			fmt.Printf("%s: %s\n", id, PrettyPrint(r))

		}
	},
}

func init() {
	serversCmd.AddCommand(serversList)
	rootCmd.AddCommand(serversCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serversCmd.PersistentFlags().StringP("name", "n", "", "Server name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serversCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
