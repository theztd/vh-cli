package dns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ztd/vh-cli/cmd/color"
	cmd "ztd/vh-cli/config"
)

func Add(zone string) {
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
	TOKEN := cmd.CFG["VH_API_KEY"]
	fmt.Println(TOKEN)
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://centrum.vas-hosting.cz/api/v1/domains/%s/dns-records", zone),
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

	//var data map[string]interface{}
	var data map[string]Record
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	for id, r := range data {
		var clr string
		switch r.Type {
		case "A":
			clr = color.Blue
		case "CNAME":
			clr = color.Green
		case "TXT":
			clr = color.Yellow

		}
		fmt.Printf("%sID: %s - %+v %s\n", clr, id, r, color.Reset)
	}

}
