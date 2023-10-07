package servers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ztd/vh-cli/config"
)

func List(server string) map[string]interface{} {
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
	TOKEN := config.CFG["VH_API_KEY"]
	URL := fmt.Sprintf("%s/servers/%s", config.CFG["VH_URL"], server)
	req, err := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", TOKEN)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	//var data map[string]Record
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data

}
