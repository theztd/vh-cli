/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package cmd

import "encoding/json"

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
