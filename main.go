package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Welcome to the cmoli.es deployment CLI!")
	showHelp()
	var choice string
	var mustPull bool
	for {
		fmt.Print(">> ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			testLocal()
			os.Exit(0)
		case "2":
			mustPull = true
			deploy(mustPull)
			os.Exit(0)
		case "3":
			mustPull = false
			deploy(mustPull)
			os.Exit(0)
		case "e":
			fmt.Println("Bye!")
			os.Exit(0)
		case "h":
			showHelp()
		default:
			fmt.Println("Invalid input")
		}
	}
}

func showHelp() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Testing local")
	fmt.Println("2. Deploy")
	fmt.Println("3. Deploy (omit git pull)")
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func testLocal() {
	cfg := newConfig()
	createContent(cfg)
	run("firefox " + filepath.Join(cfg.WebPath, "index.html"))
}

func deploy(mustPull bool) {
	fmt.Println("Starting full deployment")
	if mustPull {
		pullGitRepos()
	}
	cfg := newConfig()
	createContent(cfg)
	sendToVps(cfg)
	fmt.Println("Deployed! :)")
}

func createContent(cfg deployConfig) {
	err := prepareMdContentToConvert(cfg)
	exitIfError(err)
	convertMdToHtml(cfg.WebPath)
	err = modifyHtml(cfg)
	exitIfError(err)
	err = setMedia(cfg)
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error %v\n", err)
		os.Exit(1)
	}
}
