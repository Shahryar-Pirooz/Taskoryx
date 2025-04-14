package main

import (
	"flag"
	"os"

	"tasoryx/config"
	"tasoryx/pkg/db"
	"tasoryx/pkg/logger"
)

func main() {
	cfg := setConfig()
	logger.Init(cfg.Production)
	db.Init(cfg.Database)
	defer db.DB.Close()
}

func setConfig() config.Config {
	var path string
	const (
		defaultPath string = "./config.yml"
		message     string = "pass config to app"
	)
	flag.StringVar(&path, "c", defaultPath, message)
	flag.StringVar(&path, "config", defaultPath, message)
	flag.Parse()

	if pathEnv := os.Getenv("TASKORYX_CONFIG_PATH"); len(pathEnv) > 1 {
		path = pathEnv
	}
	cfg := config.MustRead(path)

	return cfg
}
