package dns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ztd/vh-cli/config"
)

/*
	Add(zone string, record Record) (status int, body []byte)
	Del(zone string, id string, record Record) (status int, body []byte)
	ListRecords(zone string, record Record) map[string]Record
*/

func Request(route string, method string, reqBody interface{}) (status int, data []byte) {
	jsonData, err := json.Marshal(reqBody)
	if err != nil {

		panic(err)
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", config.VH_URL, route),
		bytes.NewReader(jsonData),
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

	return res.StatusCode, body
}
