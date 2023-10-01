/*
Copyright © 2023 Marek Sirovy
*/
package main

import (
	"ztd/vh-cli/cmd"
	"ztd/vh-cli/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
