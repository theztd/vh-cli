package servers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ztd/vh-cli/cmd/color"
	"ztd/vh-cli/config"
)

func List() {
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
	// if len(zone) < 1 {
	// 	// nebyl zadan nazev zony, listuji tedy zony
	// }
	TOKEN := config.CFG["VH_API_KEY"]
	fmt.Println(TOKEN)
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://centrum.vas-hosting.cz/api/v1/servers"),
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

	for id, r := range data {
		var clr string
		// switch r.Type {
		// case "A":
		// 	clr = color.Blue
		// case "CNAME":
		// 	clr = color.Green
		// case "TXT":
		// 	clr = color.Yellow

		// }
		fmt.Printf("%sID: %s - %+v %s\n", clr, id, r, color.Reset)
	}

}
