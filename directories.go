package main

import "os"

const (
	mdPath = "/tmp/www"
)

var mediaContentPath = os.ExpandEnv("$HOME/Software/cmoli-media-content")
