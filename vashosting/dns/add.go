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

func Add(zone string, record Record) (status int, body []byte) {
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

	jsonData, err := json.Marshal(record)
	if err != nil {

		panic(err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/domains/%s/dns-records", config.VH_URL, zone),
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
