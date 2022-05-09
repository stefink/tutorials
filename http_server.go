package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye")
}

func UserAgentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Your User Agent: "+r.UserAgent())
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h2> Home Page </h2>
					<ul> 
					<li> 
					<a href="/hello">hello</a>
					</li> 
					<li>
					<a href="/ua">Get User Agent</a>
					</li>
					<li>
					<a href="/bye">Bye</a>
					</li>
					</ul>`,
	)
}
func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)
	http.HandleFunc("/ua", UserAgentHandler)
	fmt.Println("Listening on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
