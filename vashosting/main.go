package vashosting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	TOKEN   string
	URL     string
	VERSION string
)

func Get(route string, method string) map[string]interface{} {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", URL, route),
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

func Post() {

}
