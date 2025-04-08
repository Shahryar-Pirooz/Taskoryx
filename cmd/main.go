package main

import (
	"flag"
	"fmt"
	"os"

	"tasoryx/config"
)

func main() {
	fmt.Println(setConfig())
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

	fmt.Println(path)
	cfg := config.MustRead(path)

	return cfg
}
