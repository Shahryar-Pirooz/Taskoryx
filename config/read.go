package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func AbsPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)
}

func Read(path string) (Config, error) {
	var cfg Config
	fullPath, err := AbsPath(path)
	if err != nil {
		return cfg, err
	}

	viper.SetConfigFile(fullPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	return cfg, viper.Unmarshal(&cfg)
}

func MustRead(path string) Config {
	cfg, err := Read(path)
	if err != nil {
		panic(fmt.Sprintf("cannot read config path : %w", err))
	}
	return cfg
}
