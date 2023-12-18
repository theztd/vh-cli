package dns

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ztd/vh-cli/config"
)

func ListZones(zone string) map[string]interface{} {
	/*
		curl -H "X-API-Key: TOKEN" -XGET \
			https://centrum.vas-hosting.cz/api/v1/domains/example.net | jq

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
	// if len(zone) < 1 {
	// 	// nebyl zadan nazev zony, listuji tedy zony
	// }
	URL := fmt.Sprintf("%s/domains/%s", config.VH_URL, zone)
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

	if res.StatusCode != 200 {
		fmt.Println("ERR: Bad request!", URL)
		return map[string]interface{}{}
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data
}
