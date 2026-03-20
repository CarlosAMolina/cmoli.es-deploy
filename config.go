package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type deployConfig struct {
	MediaPath string `json:"media_path"`
	VpsAlias  string `json:"vps_alias"`
	WebPath   string `json:"web_path"`
}

func newConfig() deployConfig {
	configPath := filepath.Join(getCurrentPath(), "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	var config deployConfig
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config
}
