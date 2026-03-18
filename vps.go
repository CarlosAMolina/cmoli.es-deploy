package main

func sendToVps() {
	run("rsync -az --delete " + mdPath + "/ " + vpsAlias + ":" + vpsDestPath)
}
