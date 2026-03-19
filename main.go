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
	createContent()
	run("firefox " + filepath.Join(cfg.WebContentPath, "index.html"))
}

func deploy(mustPull bool) {
	fmt.Println("Starting full deployment")
	if mustPull {
		pullGitRepos()
	}
	createContent()
	sendToVps()
	fmt.Println("Deployed! :)")
}

func createContent() {
	err := prepareMdContentToConvert()
	exitIfError(err)
	convertMdToHtml(cfg.WebContentPath)
	err = modifyHtml()
	exitIfError(err)
	err = setMedia()
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error %v\n", err)
		os.Exit(1)
	}
}
