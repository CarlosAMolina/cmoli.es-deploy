package main

import "os"

const (
	mdPath      = "/tmp/www"
	vpsAlias    = "TODO"
	vpsDestPath = "~/Software/www/"
)

var mediaContentPath = os.ExpandEnv("$HOME/Software/cmoli-media-content")
