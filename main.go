/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package main

import (
	"log"
	"ztd/vh-cli/cmd"
	"ztd/vh-cli/config"
)

func main() {
	if config.DEBUG {
		log.Println("Running in debug mode...🔎")
	}
	config.Init()
	cmd.Execute()
}
