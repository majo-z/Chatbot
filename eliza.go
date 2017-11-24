//Researched and Adapted from:
//https://golang.org/pkg/net/http/
//http://www.alexedwards.net/blog/serving-static-sites-with-go
//https://golang.org/pkg/math/rand/
//https://gobyexample.com/structs
//https://golang.org/pkg/regexp/#MustCompile
//https://blog.golang.org/go-maps-in-action
//https://golang.org/pkg/bufio/#NewScanner
//https://golang.org/pkg/strings/#Replace
//https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
//https://regex101.com/
//https://data-representation.github.io/notes/regexp-go.html
//https://github.com/data-representation/eliza
//https://github.com/google/re2/wiki/Syntax

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

//+++++++++++++++++ Web server ++++++++++++++++++++++++++++++++++++++++++++++++++++

func HandleAsk(w http.ResponseWriter, r *http.Request) {
	//extracting query paramenter from Get request, it's sent from js(jquery) file
	userInput := r.URL.Query().Get("input")
	// if userInput == ""{
	// don't respond, handled in js file
	// }
	reply := Ask(userInput) //receive answer from Eliza
	fmt.Fprintf(w, reply)   //write the answer in the response writer
}

func main() {
	//go app will go to static folder and take out the index.html and serve that
	// "/" - handle any request comming in...deals with everything that starts with /
	//"./static" - look in the static folder for html file that is after the /
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//"/ask" GET/POST need parameters
	//any requests comming going to /ask are handled in HandleAsk function
	http.HandleFunc("/ask", HandleAsk)

	//listen for http file coming on to port 8080 on my computer
	http.ListenAndServe(":8080", nil)
} //main

//++++++++++++++++++++++++++++++ Eliza logic ++++++++++++++++++++++++++++++++++++++++++

var reflections map[string]string

//Struct to keep the regexp pattern and its possible answers together.
type Response struct {
	re      *regexp.Regexp
	answers []string
}

//constructor function for a Response
func NewResponse(pattern string, answers []string) Response {
	response := Response{} //create a new Response

	re := regexp.MustCompile(pattern) //create the regexp pattern
	response.re = re                  //assign the values
	response.answers = answers        //assign answers
	return response                   //return complete Response
}

//creates a list of responses by reading patterns.txt line by line
func buildResponseList() []Response {

	allResponses := []Response{} //create an empty list to hold all the responses

	file, err := os.Open("./patterns.txt")
	if err != nil { //there IS an error
		panic(err) //crash the program
	}

	//file exists!
	defer file.Close() //this will be called after this function.

	scanner := bufio.NewScanner(file) //create a scanner to read the file line by line

	for scanner.Scan() {
		//fmt.Println(scanner.Text())

		patternStr := scanner.Text()
		if strings.HasPrefix(patternStr, "#") || strings.TrimSpace(patternStr) == "" {
			continue //skip lines that are white space or start with # which are a comment
		}
		patternStr = "(?i)" + patternStr //make every pattern case insensitive
		scanner.Scan()                   //move onto the next line which holds the answers
		answersAsStr := scanner.Text()   //string that has all the answers separated by ;

		answerList := strings.Split(answersAsStr, ";") // plit into list of answers
		resp := NewResponse(patternStr, answerList)    //create a response from the pattern and list
		allResponses = append(allResponses, resp)      //add that response
	}
	return allResponses //give back all responses
} //buildResponseList

func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano()) //seed to make it return different values.
	index := rand.Intn(len(answers)) //Intn generates a number between 0 and num - 1
	return answers[index]            //can be any element
}

func reflectPronouns(original string) string {
	//reflections = readLines("file/path")// []string am:are
	if reflections == nil { // map hasn't been made yet
		reflections = map[string]string{ // will only happen once.
			"am":     "are",
			"was":    "were",
			"i":      "you",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"you":    "me",
			"me":     "you",
		}
	}
	//when we're we can be sure reflectiosn map is populated.

	//"Hello Eliza I like you"
	//{"Hello", "Eliza", "I", "Like", "you"}
	words := strings.Split(original, " ") //convert the string into a list of words

	for index, word := range words {
		//we want to change the word if it's in the map
		val, ok := reflections[word]
		if ok { //value WAS in the map
			//we want to swap with the value
			words[index] = val //eg. you -> me
		}
	}
	return strings.Join(words, " ") //convert the updated list of words back into a string
} //reflectPronouns

func Ask(userInput string) string {
	fmt.Println("Asking: " + userInput)

	responses := buildResponseList()

	for _, resp := range responses { //look at every single response/pattern/answers

		if resp.re.MatchString(userInput) {
			//extract the capture groups from the user input
			//match is a list of strings
			match := resp.re.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1] //captured is what we want to reflect pronouns on

			//remove unwanted characters like . ? and more
			chars := []string{".", "!", "?", ",", ";", ":"}
			for _, removeChar := range chars { //loop through the characters we want to replace
				//func Replace(s, old, new string, n int) string....
				//If n < 0, there is no limit on the number of replacements
				captured = strings.Replace(captured, removeChar, "", -1) //replace with empty char
			}

			captured = reflectPronouns(captured) //reflect

			formatAnswer := getRandomAnswer(resp.answers) //get random element.

			if strings.Contains(formatAnswer, "%s") { //string needs to be formatted
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer
		} //if
	} //for

	//moved default answers to patterns
	panic("We shouldn't be here. Make sure the catch all (.*) pattern exists in patterns.txt")
	//if there were no matches, one of the following is randomy picked
	//defaultAnswers := []string{"Sorry I was not listening.",
	//"Let's talk about something else, do you like music?",
	//"It's boring chating with you, tell me something interesting."}

	//return getRandomAnswer(defaultAnswers)
} //Ask
//++++++++++++++++++++++++++++++ Test ++++++++++++++++++++++++++++++++++++++++++
/*
	//patternStr := "name is (.*)" // Hello my name is bob
	//MustCompile, Compile to make a *regexp.Regexp struct
	//re := regexp.MustCompile(patternStr)

	if re.MatchString(userInput) {
		fmt.Println("There was a match!")
		//re.FindStringSubmatch()
		match := re.FindStringSubmatch(userInput)
		//match[0] is full match, match[1] is the capture group
		captured := match[1]
		fmt.Println(captured)

		formatString := "Hello %s, it's nice to meet" //this is the format string
		answer := fmt.Sprintf(formatString, captured)
		fmt.Println(answer)

	} else {
		fmt.Println("There was no match")
	}

	//slice / list of answers, and I return 1 at random
	//Hi bob
	//Hello bob
	//how's it hanging bob
*/
