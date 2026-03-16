package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func getPathSoftware() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, "Software")
}

func getCurrentPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentPath
}

func exists(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
