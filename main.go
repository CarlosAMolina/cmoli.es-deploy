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
	for {
		fmt.Print(">> ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			testLocal()
			os.Exit(0)
		case "2":
			fmt.Println("Starting full deployment")
			deploy()
			fmt.Println("Deployed! :)")
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
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func testLocal() {
	createContent()
	run("firefox " + filepath.Join(mdPath, "index.html"))
}

func deploy() {
	pullGitRepos()
	createContent()
	sendToVps()
}

func createContent() {
	err := prepareMdContentToConvert()
	exitIfError(err)
	convertMdToHtml()
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
