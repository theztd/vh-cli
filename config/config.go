package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	VERSION = "0.0.1-alfa"
	CFG     map[string]string
)

func Init() {
	if cfg, err := godotenv.Read(".vh/config.env"); err != nil {
		fmt.Println("Fatal: Unable to read config file... ", err)
		os.Exit(1)
	} else {
		CFG = cfg
	}
}
