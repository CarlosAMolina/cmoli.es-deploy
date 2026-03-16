package main

func modifyHtml() {
	path := mdPath + "/projects/rust-vs-other-languages/02-results-summary.html"
	run("sed -i 's/<table>/<table class=\"center\">/g' " + path)
}
