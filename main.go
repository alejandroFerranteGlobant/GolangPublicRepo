package main

import (
    "net/http"
	"mux"
)

var logActivity = true

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/assets/", handleUpload).Methods(http.MethodPost)
	router.HandleFunc("/assets/", flushDownloadDirectory).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", router)	
	
}




