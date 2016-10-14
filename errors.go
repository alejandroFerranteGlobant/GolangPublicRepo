package main

import(
    "net/http"
    "fmt"
)

func emptyURLError(w http.ResponseWriter){
    fmt.Fprintln(w, "ERROR: The URL is empty")
}


func failedFileCreationError(){

}
    
func    failedCoyingFileError(){

}