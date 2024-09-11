/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package dns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ztd/vh-cli/config"
)

func Del(zone string, id string, record Record) (status int, body []byte) {
	/*
		curl -H "X-API-Key: TOKEN" -XPOST \
			'https://centrum.vas-hosting.cz/api/v1/domains/<domain>/dns-records' \
			--header 'Content-Type: application/json' \
			--data '{
						"name": "test",
						"content": "9.8.7.0",
						"type": "A",
						"ttl": 1800
					}
					'

		>>>
	*/
	// if len(zone) < 1 {
	// 	// nebyl zadan nazev zony, listuji tedy zony
	// }
	jsonData, err := json.Marshal(record)
	if err != nil {

		panic(err)
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/domains/%s/dns-records/%s", config.VH_URL, zone, id),
		bytes.NewReader(jsonData),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", config.VH_API_KEY)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERR [request]:", err)
		//panic(err)
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return res.StatusCode, body
}
