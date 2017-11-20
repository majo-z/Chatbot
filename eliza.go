package main

import (
	"net/http"
)

func main() {

	//go app will go to static folder and take out the index.html and serve that
	// "/" - handle any request comming in...deals with everything that starts with /
	//"./static" - look in the static folder for html file that is after the /
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//listen for http file coming on to port 8080 on my computer
	http.ListenAndServe(":8080", nil)

} //main
