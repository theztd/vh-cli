/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"ztd/vh-cli/vashosting/servers"

	"github.com/spf13/cobra"
)

// serversCmd represents the servers command
var serversCmd = &cobra.Command{
	Use:     "servers",
	Aliases: []string{"server", "srv", "s"},
	Short:   "Sprava serveru",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

var serversListJson = &cobra.Command{
	Use:     "list-json",
	Aliases: []string{"lj"},
	Short:   "Vypise seznam serveru jako RAW json (tak jak ho dostane od API VH)",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		for id, r := range servers.ListJson(name) {
			//fmt.Printf("%s - %+v\n", id, r)
			// var clr string
			// status := r["status"].(string)
			// switch status {
			// case "A":
			// 	clr = color.Blue
			// case "CNAME":
			// 	clr = color.Green
			// case "TXT":
			// 	clr = color.Yellow

			// }
			// fmt.Printf("%sID: %s - %+v %s\n", clr, id, r, color.Reset)

			fmt.Printf("%s: %s\n", id, PrettyPrint(r))

		}
	},
}

var serversList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Vypise seznam serveru",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		tmplf, _ := cmd.Flags().GetString("template-file")
		grpBy, _ := cmd.Flags().GetString("group-by")
		filterLabels, _ := cmd.Flags().GetStringSlice("filter-labels")

		if tmplf != "" {
			tmplContent, err := os.ReadFile(tmplf) // #nosec G304
			if err != nil {
				fmt.Println("Unable to read template file", err)
				os.Exit(1)
			}
			if grpBy != "" {
				res := servers.GroupResultsBy(
					servers.FilterServerByLabel(servers.List(name), filterLabels), grpBy)
				RenderTemplate(string(tmplContent), res, os.Stdout) // #nosec G104
			} else {
				RenderTemplate(string(tmplContent), servers.List(name), os.Stdout) // #nosec G104
			}

			os.Exit(0)
		}

		// Vystup bez sablony s groupovanim
		if grpBy != "" {
			res := servers.GroupResultsBy(
				servers.FilterServerByLabel(servers.List(name), filterLabels), grpBy)
			for key, servers := range res {
				fmt.Printf("#######  %s   #######\n", key)
				for _, s := range servers {
					fmt.Println(s.ID, s.DisplayName, s.Name, s.HW, s.Status, s.Ram, s.IPv4, s.IPv6, s.Os, s.Storage, s.Labels)
				}

			}
			os.Exit(0)
		}

		res := servers.ListJson(name)
		for id, r := range res {
			fmt.Printf("%s: %s\n", id, PrettyPrint(r))
		}
	},
}

func init() {
	serversCmd.AddCommand(serversList)
	serversCmd.AddCommand(serversListJson)
	rootCmd.AddCommand(serversCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serversCmd.PersistentFlags().StringP("name", "n", "", "Server name")
	serversCmd.PersistentFlags().StringP("template-file", "f", "", "Template file path")
	serversCmd.PersistentFlags().StringP("group-by", "G", "", "Group results by")
	serversCmd.PersistentFlags().StringSliceP("filter-labels", "", []string{}, "Filter results by label")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serversCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
