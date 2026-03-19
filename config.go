package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var cfg = loadDeployConfig()

type deployConfig struct {
	MediaContentPath string `json:"media_content_path"`
	VpsAlias         string `json:"vps_alias"`
	VpsDestPath      string `json:"vps_dest_path"`
	WebContentPath           string `json:"web_content_path"`
}

func loadDeployConfig() deployConfig {
	configPath := filepath.Join(getCurrentPath(), "config-deploy.json")
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
