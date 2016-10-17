package main

import (
    "net/http"
	"mux" 
)

var logActivity = true

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/assets/save", handleUpload).Methods("POST")
	http.ListenAndServe(":8080", router)	

}