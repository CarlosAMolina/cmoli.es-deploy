package main

// Connect: docker exec -it python-create-pandoc-script-container /bin/sh
func modifyHtmlDocker() {
	path := getVolumePath("nginx-web-content") + "/projects/rust-vs-other-languages/02-results-summary.html"
	run("sed -i 's/<table>/<table class=\"center\">/g' " + path)
}

func modifyHtml() {
	path := mdPath + "/projects/rust-vs-other-languages/02-results-summary.html"
	run("sed -i 's/<table>/<table class=\"center\">/g' " + path)
}
