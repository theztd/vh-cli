/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package cmd

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

func RenderTemplate(tmplContent string, data any, out io.Writer) (err error) {
	/*
		Custom template rendering including template functions

	*/
	tmplOut := template.New("").Funcs(template.FuncMap{
		"Contains": func(data []string, search string) bool {
			for _, s := range data {
				if s == search {
					return true
				}
			}
			return false
		},
		"Replace":    strings.Replace,
		"ReplaceAll": strings.ReplaceAll,
		"Join":       strings.Join,
	})
	tmplOut, err = tmplOut.Parse(tmplContent)
	if err != nil {
		fmt.Println("Unable to parse tempate", err)
		return err
	}

	return tmplOut.Execute(out, data) // #nosec G104
}
