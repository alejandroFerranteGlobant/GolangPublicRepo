package main

import (
    "fmt"
    "net/http"
	"strings"
)

var LOG_ACTIVITY = true

func handler(w http.ResponseWriter, r *http.Request) {
	
	if(LOG_ACTIVITY){
		fmt.Println("Reqeuest Method: ",r.Method)
		fmt.Println("Reqeuest Header: ",r.Header)
		fmt.Println("URL Recieved: ",r.URL)
		fmt.Println("URL Mapping: ",strings.Split(r.URL.String(),"/"))
	}
	
    
	switch method := r.Method; method {
		case "POST":
			handlePost(w,r)
		case "GET":
			handleGet(w,r)
	}

	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}