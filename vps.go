package main

func sendToVps(cfg deployConfig) {
	run("rsync -az --delete " + cfg.WebContentPath + "/ " + cfg.VpsAlias + ":" + cfg.VpsDestPath)
}
