package main

import (
	"fmt"
	"net/http"
)

func Ask(input string) string {
	//do regex
	//read from file
	//process input
	//sub extracted words into answer
	//return answer
	return input
}

func HandleAsk(w http.ResponseWriter, r *http.Request) {
	// r.URL == "http:/localhost:8080/ask?input=hello"
	userInput := r.URL.Query().Get("input") // extracts hello
	// if userInput == ""{
	// don't respond, handled in js file
	// }
	reply := Ask(userInput)
	fmt.Fprintf(w, reply)

}

func main() {

	//go app will go to static folder and take out the index.html and serve that
	// "/" - handle any request comming in...deals with everything that starts with /
	//"./static" - look in the static folder for html file that is after the /
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//"/ask" GET/POST need parameters
	http.HandleFunc("/ask", HandleAsk)

	//listen for http file coming on to port 8080 on my computer
	http.ListenAndServe(":8080", nil)

} //main
