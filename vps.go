package main

func sendToVps() {
	run("rsync -az --delete " + cfg.MdPath + "/ " + cfg.VpsAlias + ":" + cfg.VpsDestPath)
}
