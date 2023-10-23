/*
Copyright © 2023 Marek Sirovy
*/
package main

import (
	"log"
	"ztd/vh-cli/cmd"
	"ztd/vh-cli/config"
)

func main() {
	config.Init()
	if config.DEBUG {
		log.Println("Running in debug mode...🔎")
	}
	cmd.Execute()
}
