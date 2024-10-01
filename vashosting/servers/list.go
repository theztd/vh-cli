/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package servers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"ztd/vh-cli/config"
)

type ServerOut struct {
	ID          uint32 `json:"id"`
	Name        string
	DisplayName string                 `json:"displayName"`
	Os          string                 `json:"operatingSystem"`
	Addresses   map[string]interface{} `json:"addresses"`
	IPv4        []string
	IPv6        []string
	Labels      []string          `json:"labels"`
	Storage     map[string]uint32 `json:"storage"`
	Ram         int               `json:"ram"`
	Status      string            `json:"status"`
	HW          string            `json:"dedicatedServerName"`
}

func List(server string) map[string]ServerOut {
	/*
		curl -H "X-API-Key: TOKEN" -XGET \
			https://centrum.vas-hosting.cz/api/v1/servers | jq

		>>>
		{
			"127898": {
				"name": "example.net",
				"type": "NS",
				"content": "ns.example.cz",
				"priority": null,
				"ttl": 1800,
				"created": "2019-11-04 11:51:40",
				"updated": "2025-08-13 15:04:23"
			},
			...
		}
	*/
	URL := fmt.Sprintf("%s/servers/%s", config.VH_URL, server)
	req, err := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", config.VH_API_KEY)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]ServerOut
	ret := map[string]ServerOut{}
	if err := json.Unmarshal(body, &data); err != nil {
		if config.DEBUG {
			fmt.Printf("\nDEBUG [vashosting.servers.List]: %s\n\n", body)
			panic(err)
		}
		fmt.Printf("Broken response from API: %s\n", body)
		os.Exit(1)

	}

	for name, d := range data {
		tmp := d
		tmp.Name = name
		for ip4, _ := range d.Addresses["ipv4"].(map[string]interface{}) {
			tmp.IPv4 = append(tmp.IPv4, ip4)
		}

		ipv6map, ok := d.Addresses["ipv6"].(map[string]interface{})
		// fmt.Println(ipv6map)
		if !ok {
			// skip broken inputs
			continue
		}
		for ip6, _ := range ipv6map {
			tmp.IPv6 = append(tmp.IPv6, ip6)
		}
		ret[name] = tmp

	}
	return ret

}

func ListJson(server string) map[string]interface{} {
	/*
		curl -H "X-API-Key: TOKEN" -XGET \
			https://centrum.vas-hosting.cz/api/v1/servers | jq

		>>>
		{
			"127898": {
				"name": "example.net",
				"type": "NS",
				"content": "ns.example.cz",
				"priority": null,
				"ttl": 1800,
				"created": "2019-11-04 11:51:40",
				"updated": "2025-08-13 15:04:23"
			},
			...
		}
	*/
	URL := fmt.Sprintf("%s/servers/%s", config.VH_URL, server)
	req, err := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", config.VH_API_KEY)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data

}

func GroupResultsBy(data map[string]ServerOut, grpBy string) map[string][]ServerOut {
	/*
		Group results by key
	*/

	// Find correct group key name (I want to be able accept case insensitive parameter)
	var groupByKey string
	structKeys := reflect.TypeOf(ServerOut{})
	for i := 0; i < structKeys.NumField(); i++ {
		// compare strings case insensitive
		if strings.EqualFold(structKeys.Field(i).Name, grpBy) {
			groupByKey = structKeys.Field(i).Name
		}
	}

	// When groupByKey is stay empty, the key is not in struct...
	if groupByKey == "" {
		fmt.Printf("Zadano nespravne, nebo neexistujici pole '%s'! Zkus nektery z existujicich klicu:\n", grpBy)
		for i := 0; i < structKeys.NumField(); i++ {
			f := structKeys.Field(i)
			if f.Type.Kind() == reflect.String {
				fmt.Println(" - ", f.Name)
			}
		}

		return nil
	} else {
		if config.DEBUG {
			fmt.Println("DEBUG [GroupResultsBy]: Groupuju dle klice", groupByKey)
		}
	}

	newRet := map[string][]ServerOut{}
	for _, srv := range data {
		key := reflect.ValueOf(srv).FieldByName(groupByKey)
		newRet[key.String()] = append(newRet[key.String()], srv)

	}
	return newRet
}

func FilterServerByLabel(data map[string]ServerOut, filterLabels []string) (ret map[string]ServerOut) {
	/*
		Filter results by name and kind
	*/
	if config.DEBUG {
		log.Println("DEBUG [FilterServerByLabel]: Filter by labels", filterLabels)
	}
	ret = map[string]ServerOut{}
	for id, d := range data {
		matchScore := 0

		// Iterate over servr labels and check if given filterLabels match all of them
		for _, srvlabel := range d.Labels {
			for _, testLabel := range filterLabels {
				if strings.Contains(testLabel, srvlabel) {
					matchScore++
				}
			}
		}

		// Compare if amount of matchScore is same to number of items in the Filter structure
		if matchScore == len(filterLabels) {
			if config.DEBUG {
				log.Println("DEBUG [FilterServerByLabel]: Server", d.Name, "with labels ", d.Labels, "has been appended because contains required labels.")
			}
			ret[string(id)] = d
		} else {
			if config.DEBUG {
				log.Println("DEBUG [FilterServerByLabel]: Skip server", d.Name)
			}
		}
	}
	return ret
}
