package main

import (
	"os"
	"strings"
)

func modifyHtml() error {
	path := cfg.WebContentPath + "/projects/rust-vs-other-languages/02-results-summary.html"
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	modified := strings.ReplaceAll(string(content), "<table>", "<table class=\"center\">")
	return os.WriteFile(path, []byte(modified), 0644)
}
