package main

import (
	"fmt"
	"os"
	"os/user"
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
			fmt.Println("Starting full deployment")
			deploy()
			fmt.Println("Deployed! :)")
			os.Exit(0)
		case "2":
			testLocal()
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
	fmt.Println("1. Deploy")
	fmt.Println("2. Testing local")
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func deploy() {
	pullGitRepos()
	err := prepareMdContentToConvert()
	exitIfError(err)
	convertMdToHtml()
	modifyHtml()
	copyMedia()
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error %v\n", err)
		os.Exit(1)
	}
}

func testLocal() {
	deploy()
	run("firefox " + filepath.Join(mdPath, "index.html"))
}

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

func copyMedia() {
	run("cp -r ~/Software/cmoli-media-content/* " + mdPath)
	videoVolumePath := filepath.Join(mdPath, "felices-fiestas/src/movie.mp4")
	run("rm " + videoVolumePath)
	run("ln -s ~/Software/cmoli-media-content/felices-fiestas/src/movie.mp4 " + videoVolumePath)
}
