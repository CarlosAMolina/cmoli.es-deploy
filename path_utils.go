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
		if filepath.Base(srcPath) == "favicon.ico" {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				fmt.Printf("read error: srcPath=%s err=%v\n", srcPath, err)
				return err
			}
			err = os.WriteFile(dstPath, data, info.Mode())
			if err != nil {
				fmt.Printf("write error: dstPath=%s err=%v\n", dstPath, err)
				return err
			}
			fmt.Printf("favicon copied: %s -> %s\n", srcPath, dstPath)
			return nil
		} else {
			fmt.Printf("symlink created: %s -> %s\n", srcPath, dstPath)
			return os.Symlink(srcPath, dstPath)
		}
	})
}

// TODO improve, modity to only remove symlinks, rename to removeSymlinks and drop if condition of favicon.ico
func removeMedia(cfg deployConfig) error {
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
		if filepath.Base(srcPath) == "favicon.ico" {
			return nil
		}
		relPath, err := filepath.Rel(cfg.MediaPath, srcPath)
		if err != nil {
			fmt.Printf("rel error: srcRoot=%s srcPath=%s err=%v\n", cfg.MediaPath, srcPath, err)
			return err
		}
		dstPath := filepath.Join(cfg.WebPath, relPath)
		fmt.Printf("symlink removed: %s\n", dstPath)
		return os.Remove(dstPath)
	})
}
