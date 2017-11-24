# Eliza Chatbot Project 
My name is Marian and I am 3rd year software development student at GMIT. 
Objective of this project was to create a chatbot web application to demonstrate the superficiality of communication between man and machine.
You can find more inforamtion about Eliza project created between 1964-1966 at https://en.wikipedia.org/wiki/ELIZA.
The programs has be written in Go programming language. 

## How to download and run the chatbot 
The repository can be found at https://github.com/majo-z/Chatbot.git. 
To run the code in this repository, the files must first be compiled. The [Go compiler](https://golang.org/dl/) and [Git](https://git-scm.com/) must first be installed on your machine.
Once they are installed, the code can be compiled and run by following the steps below.
We assume you are using the command line.

1. Open the Git Bash and clone the repository 

```bash
> git clone https://github.com/majo-z/Chatbot.git
```
2. Change into the folder.
```bash
> cd Chatbot
```
3. Compile the eliza.go file with the following command.
```bash
> go build eliza.go
```
4. Run the executable produced.
```bash
> ./eliza.exe
eliza.exe is now running on the background
```
5. Navigate to your browser and ype in the address bar
```bash
> localhost: 8080
or
> 127.0.0.1:8080
```
6. To close the program
```bash
> exit the browser
>in the command line crtl+c 
or
>in the task manager processes find eliza.exe and end task
```

## References
https://golang.org/pkg/net/http/
http://www.alexedwards.net/blog/serving-static-sites-with-go
https://golang.org/pkg/math/rand/
https://gobyexample.com/structs
https://golang.org/pkg/regexp/#MustCompile
https://blog.golang.org/go-maps-in-action
https://golang.org/pkg/bufio/#NewScanner
https://golang.org/pkg/strings/#Replace
https://github.com/data-representation/go-ajax/blob/master/webapp.go
https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
https://golang.org/doc/articles/wiki/
https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
https://regex101.com/
https://data-representation.github.io/notes/regexp-go.html
https://github.com/data-representation/eliza/blob/master/data/responses.txt
https://github.com/data-representation/eliza
https://github.com/google/re2/wiki/Syntax
https://data-representation.github.io/notes/ajax.html
https://api.jquery.com/keypress/
https://api.jquery.com/event.preventdefault/
https://msdn.microsoft.com/en-us/library/mt260494.aspx
https://getbootstrap.com/docs/4.0/getting-started/introduction/
https://stackoverflow.com/questions/11715646/scroll-automatically-to-the-bottom-of-the-page
https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
https://thoughtcatalog.com/oliver-miller/2012/08/a-conversation-with-eliza/
https://github.com/google/re2/wiki/Syntax
https://github.com/data-representation/eliza/blob/master/data/responses.txt

## Learn more about Go language
https://golang.org/
https://golang.org/doc/code.html
https://play.golang.org/
