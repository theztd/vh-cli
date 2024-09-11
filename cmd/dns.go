/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"ztd/vh-cli/vashosting/dns"

	"github.com/spf13/cobra"
)

var dnsListRecords = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Vypis zaznamu v domene",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		outFmt, _ := cmd.Flags().GetString("out")
		zone, _ := cmd.Flags().GetString("zone")
		name, _ := cmd.Flags().GetString("name")
		kind, _ := cmd.Flags().GetString("type")
		ttl, _ := cmd.Flags().GetInt("ttl")
		value, _ := cmd.Flags().GetString("value")
		tmplf, _ := cmd.Flags().GetString("template-file")
		rec := dns.Record{
			Name:    name,
			Type:    kind,
			TTL:     ttl,
			Content: value,
		}

		if tmplf != "" {
			tmplContent, err := os.ReadFile(tmplf) // #nosec G304
			if err != nil {
				fmt.Println("Unable to read template file", err)
				os.Exit(1)
			}

			RenderTemplate(string(tmplContent), dns.ListRecords(zone, rec, dns.Filter{Kind: kind, Name: name}), os.Stdout) // #nosec G104
			os.Exit(0)
			// don't continue
		}

		for id, r := range dns.ListRecords(zone, rec, dns.Filter{Kind: kind, Name: name}) {
			if outFmt == "csv" {
				fmt.Printf("%s;%s;%s;%d;%s\n", id, r.Name, r.Type, r.TTL, r.Content)
			} else {
				fmt.Printf("%s: %s\n", id, PrettyPrint(r))
			}

		}
	},
}

var dnsAddRecord = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Prida zaznam do domeny",
	Long: `Zapis neni okamzity!!!

Novy zaznam v DNS se muze objevit az po 15minutach
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//outFmt, _ := cmd.Flags().GetString("out")
		zone, _ := cmd.Flags().GetString("zone")
		name, _ := cmd.Flags().GetString("name")
		kind, _ := cmd.Flags().GetString("type")
		ttl, _ := cmd.Flags().GetInt("ttl")
		value, _ := cmd.Flags().GetString("value")
		rec := dns.Record{
			Name:    name,
			Type:    kind,
			TTL:     ttl,
			Content: value,
		}

		status, body := dns.Add(zone, rec)
		if status != 204 {
			fmt.Println("ERR: StatusCode:", status)
			fmt.Println("ERR: Response msg:", string(body))
		}
	},
}

var dnsDelRecord = &cobra.Command{
	Use:     "del",
	Aliases: []string{"rm"},
	Short:   "Smaze zaznam z domeny",
	Long: `Zmena neni okamzita!!!

Novy zaznam z DNS muze zmizet az po 15minutach
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//outFmt, _ := cmd.Flags().GetString("out")
		zone, _ := cmd.Flags().GetString("zone")
		name, _ := cmd.Flags().GetString("name")
		kind, _ := cmd.Flags().GetString("type")
		ttl, _ := cmd.Flags().GetInt("ttl")
		value, _ := cmd.Flags().GetString("value")
		rec := dns.Record{
			Name:    name,
			Type:    kind,
			TTL:     ttl,
			Content: value,
		}

		recordsForDeleting := dns.ListRecords(zone, rec, dns.Filter{Kind: kind, Name: name})

		if len(recordsForDeleting) == 0 {
			fmt.Println("Neni co mazat.")
			return
		}

		fmt.Println("Chystate se smazat nasledujici zaznamy:")
		for id, r := range recordsForDeleting {
			fmt.Println(id, r)
		}
		buf := bufio.NewReader(os.Stdin)
		fmt.Printf("Jste si tim jisty?\n(yes/no)> ")
		confirm, _ := buf.ReadString('\n')
		if confirm == "yes\n" {
			for id, r := range recordsForDeleting {
				status, body := dns.Del(zone, id, r)
				if status != 204 {
					fmt.Println("ERR: StatusCode:", status)
					fmt.Println("ERR: Response msg:", string(body))
				}
			}

		} else {
			fmt.Println("Nic nemazu")
		}
	},
}

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Sprava DNS zaznamu",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

func init() {
	dnsCmd.AddCommand(dnsListRecords, dnsAddRecord, dnsDelRecord)
	rootCmd.AddCommand(dnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	dnsCmd.PersistentFlags().StringP("zone", "z", "", "DNS zone name")
	dnsCmd.PersistentFlags().StringP("out", "o", "", "Output format")
	dnsCmd.PersistentFlags().StringP("name", "n", "", "DNS record name")
	dnsCmd.PersistentFlags().StringP("type", "t", "TXT", "DNS record type")
	dnsCmd.PersistentFlags().StringP("value", "v", "", "DNS record value")
	dnsCmd.PersistentFlags().IntP("prio", "p", 0, "DNS record priority")
	dnsCmd.PersistentFlags().IntP("ttl", "T", 86400, "DNS record TTL")
	dnsCmd.PersistentFlags().StringP("template-file", "f", "", "Template file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
