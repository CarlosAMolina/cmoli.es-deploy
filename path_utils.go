package main

import (
	"fmt"
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
	return createSymlinksRecursive(cfg.MediaContentPath, cfg.WebContentPath)
}

func createSymlinksRecursive(srcRoot, dstRoot string) error {
	return filepath.Walk(srcRoot, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("walk error: srcRoot=%s srcPath=%s err=%v\n", srcRoot, srcPath, err)
			return err
		}
		if info == nil {
			fmt.Printf("walk nil info: srcRoot=%s srcPath=%s\n", srcRoot, srcPath)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			fmt.Printf("rel error: srcRoot=%s srcPath=%s err=%v\n", srcRoot, srcPath, err)
			return err
		}
		dstPath := filepath.Join(dstRoot, relPath)
		fmt.Printf("symlink %s -> %s\n", srcPath, dstPath)
		return os.Symlink(srcPath, dstPath)
	})
}
