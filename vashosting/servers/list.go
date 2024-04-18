package servers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ztd/vh-cli/config"
)

type ServerOut struct {
	ID        uint32 `json:"id"`
	Name      string
	Os        string                 `json:"operatingSystem"`
	Addresses map[string]interface{} `json:"addresses"`
	IPv4      []string
	IPv6      []string
	Labels    []string          `json:"labels"`
	Storage   map[string]uint32 `json:"storage"`
	Ram       int               `json:"ram"`
	Status    string            `json:"status"`
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
	//var data map[string]Record
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
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
			fmt.Println(ok)
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
	//var data map[string]Record
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data

}
