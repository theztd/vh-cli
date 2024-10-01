/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package main

import (
	"fmt"
	"log"
	"ztd/vh-cli/cmd"
	"ztd/vh-cli/config"
)

func main() {
	if config.DEBUG {
		log.Println("Running in debug mode...🔎")
	}
	err := config.Init()
	if err != nil {
		fmt.Println("FATAL: Unable to init config.", err)
	}
	cmd.Execute()
}
