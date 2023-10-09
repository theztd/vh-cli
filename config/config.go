package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	VERSION = "0.0.1-alfa"
	// CFG     map[string]string
	VH_API_KEY  string
	VH_URL      string
	DEFAULT_TTL int
)

func GetEnv(val string, def string) string {
	if len(val) != 0 {
		return val
	}
	return def
}

func Init() {
	// Nejdriv se zkusi vzit config z lokalni cesty
	// a pokud neexistuje, vezme se ten z home
	cfgFile := GetEnv(".vh/config.env", "~/.vh/config.env")
	if cfg, err := godotenv.Read(cfgFile); err != nil {
		fmt.Println("Fatal: Unable to read config file... ", cfgFile, err)
		os.Exit(1)
	} else {
		// CFG = cfg
		VH_API_KEY = GetEnv(cfg["VH_API_KEY"], "")
		VH_URL = GetEnv(cfg["VH_URL"], "")
		if defttl, err := strconv.Atoi(GetEnv(cfg["DEFAULT_TTL"], "86400")); err != nil {
			fmt.Println("ERR [config]: Wrong config format DEFAULT_TTL")
			panic(err)
		} else {
			DEFAULT_TTL = defttl
		}
	}
}
