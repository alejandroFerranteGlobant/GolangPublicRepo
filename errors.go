package main

import(
    //"net/http"
    "fmt"
    "net/http"
)


func multipartParsingError(w http.ResponseWriter, r *http.Request, theError error){
    if(logActivity){fmt.Println("Multipart Parsing Error: ",theError.Error())}
}

func fileOpenError(w http.ResponseWriter, r *http.Request, theError error){
    if(logActivity){fmt.Println("File Opening Error: ",theError.Error())}
}

func fileCreationError(w http.ResponseWriter, r *http.Request, theError error){
    if(logActivity){fmt.Println("File Creation Error: ",theError.Error())}
}