package main

func sendToVps(cfg deployConfig) {
	run("rsync -az --delete " + cfg.WebPath + "/ " + cfg.VpsAlias + ":" + cfg.WebPath)
}
