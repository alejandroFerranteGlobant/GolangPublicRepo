package main

import (
    "net/http"
    "fmt"
    "io"
    "os"
)

func handleUpload(w http.ResponseWriter, r *http.Request){

	doUpload(w,r,"C:/TST/")//TODO: Make this dynamic; pass path as parameter int the request or URL
	w.WriteHeader(http.StatusOK)

}


func doUpload(w http.ResponseWriter, r *http.Request, destinationPath string){

    //PARSE MULTIPART
    parsingError := r.ParseMultipartForm(2000000)
    if(parsingError != nil){
        multipartParsingError(w,r)
        return
    }

    //GET MULIPART FILE MAP
    formdata := r.MultipartForm
    if(logActivity){fmt.Println("MAP: ",formdata.File,":",len(formdata.File))}
    
    //ITERATE ALL FILES
    i := 1
    for _, fileHandlers := range formdata.File { 
        for _, fileHandler := range fileHandlers {
            
            //CREATE DESTINATION FILE ON DISK
            dst, creationError := os.Create( fmt.Sprintf("%s/RecievedFile%d.txt", destinationPath,i) )
            if(creationError != nil){
                fileCreationError(w,r)
                return
            }
            i++

            //OPEN FILE             
            file, err := fileHandler.Open()
            if err != nil {
 			    fileOpenError(w,r)
 			    return
 		    }
            
            //COPY FILE TO DESTINATION ON DISK
            io.Copy(dst, file);
 		    defer file.Close()
 		    
        }

 	}
    
}

