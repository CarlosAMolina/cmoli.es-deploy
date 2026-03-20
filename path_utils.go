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

func removeContents(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			if err := removeContents(fullPath); err != nil {
				return err
			}
			if err := os.Remove(fullPath); err != nil {
				return err
			}
		} else {
			if err := os.Remove(fullPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func setMedia(cfg deployConfig) error {
	return filepath.Walk(cfg.MediaPath, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("walk error: srcRoot=%s srcPath=%s err=%v\n", cfg.MediaPath, srcPath, err)
			return err
		}
		if info == nil {
			fmt.Printf("walk nil info: srcRoot=%s srcPath=%s\n", cfg.MediaPath, srcPath)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(cfg.MediaPath, srcPath)
		if err != nil {
			fmt.Printf("rel error: srcRoot=%s srcPath=%s err=%v\n", cfg.MediaPath, srcPath, err)
			return err
		}
		dstPath := filepath.Join(cfg.WebPath, relPath)
		fmt.Printf("symlink %s -> %s\n", srcPath, dstPath)
		return os.Symlink(srcPath, dstPath)
	})
}
