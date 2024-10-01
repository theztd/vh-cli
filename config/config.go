/*
Copyright © 2023 Marek Sirovy msirovy@gmail.com
*/
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

func Init() error {
	// Nejdriv se zkusi vzit config z lokalni cesty
	// a pokud neexistuje, vezme se ten z home
	cfgFile := GetEnv(os.Getenv("VH_CONFIG_PATH"), "~/.vh/config.env")

	// Expand ~ to full path
	if strings.HasPrefix(cfgFile, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("ERR [config.Init]: Unable to expand ~ to full path.")
		}
		cfgFile = filepath.Join(home, cfgFile[1:])
	}

	// Read env variables from cfgFile
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

		return nil

	} else {
		fmt.Println("Fatal: Unable to read config file... ", cfgFile, err)
		return err
	}
}
