/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"
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

var serversList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Vypise seznam serveru",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		tmplf, _ := cmd.Flags().GetString("template-file")
		if tmplf != "" {
			tmplContent, err := os.ReadFile(tmplf) // #nosec G304
			if err != nil {
				fmt.Println("Unable to read template file", err)
				os.Exit(1)
			}

			tmplOut := template.New("").Funcs(template.FuncMap{
				"Contains": func(data []string, search string) bool {
					for _, s := range data {
						if s == search {
							return true
						}
					}
					return false
				},
				"Replace": func(data string, before string, after string) string {
					return strings.ReplaceAll(data, before, after)
				},
			})
			tmplOut, err = tmplOut.Parse(string(tmplContent))
			if err != nil {
				fmt.Println("Unable to parse tempate", err)
				os.Exit(1)
			}

			tmplOut.Execute(os.Stdout, servers.List(name)) // #nosec G104
			os.Exit(0)
		}

		// clr := ""
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

func init() {
	serversCmd.AddCommand(serversList)
	rootCmd.AddCommand(serversCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serversCmd.PersistentFlags().StringP("name", "n", "", "Server name")
	serversCmd.PersistentFlags().StringP("template-file", "f", "", "Template file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serversCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
