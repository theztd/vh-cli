/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// CFG     map[string]string
	VH_API_KEY  string
	VH_URL      string
	DEFAULT_TTL int
	DEBUG       bool
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
	if cfg, err := godotenv.Read(cfgFile); err == nil {
		// Pokud je konfigurace v ENV, pouzij tu, jinak zkus config, pripadne pouzij default
		VH_URL = GetEnv(os.Getenv("VH_URL"), GetEnv(cfg["VH_URL"], ""))
		DEBUG, _ = strconv.ParseBool(GetEnv(os.Getenv("DEBUG"), GetEnv(cfg["DEBUG"], "false")))
		if defttl, err := strconv.Atoi(GetEnv(os.Getenv("DEFAULT_TTL"), GetEnv(cfg["DEFAULT_TTL"], "86400"))); err != nil {
			fmt.Println("ERR [config]: Wrong config format DEFAULT_TTL")
			panic(err)
		} else {
			DEFAULT_TTL = defttl
		}
		// API key nema default
		VH_API_KEY = GetEnv(os.Getenv("VH_API_KEY"), cfg["VH_API_KEY"])

	} else {
		fmt.Println("Fatal: Unable to read config file... ", cfgFile, err)
		os.Exit(1)
	}
}
