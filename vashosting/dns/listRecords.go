/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package dns

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"ztd/vh-cli/config"
)

type Record struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
}

// Used only for filtering results
type Filter struct {
	Kind string
	Name string
}

func filterDnsResults(data map[string]Record, filter Filter) (ret map[string]Record) {
	/*
		Filter results by name and kind
	*/
	ret = map[string]Record{}
	for id, d := range data {
		matchScore := 0
		if filter.Kind == d.Type || filter.Kind == "" {
			matchScore++
		}
		if strings.Contains(d.Name, filter.Name) || filter.Name == "" {
			matchScore++
		}

		// Compare if amount of matchScore is same to number of items in the Filter structure
		if matchScore == reflect.ValueOf(Filter{}).NumField() {
			ret[string(id)] = d
		}
	}
	return ret
}

func ListRecords(zone string, record Record, filter Filter) (ret map[string]Record) {
	if config.DEBUG {
		fmt.Println("Filter results by:", filter)
	}

	status, data := Request(
		fmt.Sprintf("domains/%s/dns-records", zone),
		"GET",
		nil,
	)

	if status != 200 {
		log.Fatalln("[ListRecords] Response code is", status)
		if config.DEBUG {
			log.Println(string(data))
		}

	}

	if err := json.Unmarshal(data, &ret); err != nil {
		log.Println("[ListRecords] Unable to parse response")
		if config.DEBUG {
			log.Println(string(data))
		}
		return nil
	}

	return filterDnsResults(ret, filter)
}
