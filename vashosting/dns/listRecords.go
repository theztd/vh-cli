package dns

import (
	"encoding/json"
	"fmt"
	"log"
	"ztd/vh-cli/config"
)

type Record struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
}

func ListRecords(zone string, record Record) (ret map[string]Record) {
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

	return ret

}
