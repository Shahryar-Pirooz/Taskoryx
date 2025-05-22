package main

import (
	"flag"
	"os"

	"tasoryx/api/http"
	"tasoryx/app"
	"tasoryx/config"
	"tasoryx/pkg/logger"
)

func main() {
	cfg := setConfig()
	logger.Init(false)
	app := app.NewApp(cfg)
	if err := http.Run(app, cfg.Server); err != nil {
		logger.Get().Error(err.Error())
	}
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
