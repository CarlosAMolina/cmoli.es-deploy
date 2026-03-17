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

func setMedia() error {
	return createSymlinksRecursive(mediaContentPath, mdPath)
}

func createSymlinksRecursive(srcRoot, dstRoot string) error {
	return filepath.Walk(srcRoot, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstRoot, relPath)
		return os.Symlink(srcPath, dstPath)
	})
}
