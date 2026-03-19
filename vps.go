package main

func sendToVps() {
	run("rsync -az --delete " + cfg.WebContentPath + "/ " + cfg.VpsAlias + ":" + cfg.VpsDestPath)
}
