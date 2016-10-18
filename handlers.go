package main

import (
    "net/http"
    "fmt"
    "io"
    "os"
)

func handleUpload(w http.ResponseWriter, r *http.Request){

	doUpload(w,r)
	w.WriteHeader(http.StatusOK)

}


func doUpload(w http.ResponseWriter, r *http.Request){

    //LOAD SETTINGS
    sysSettings := GetSettings()

    //PARSE MULTIPART
    parsingError := r.ParseMultipartForm(sysSettings.ParsingMemoryLimit)
    if(parsingError != nil){
        multipartParsingError(w,r,parsingError)
        return
    }

    //GET MULIPART FILE MAP
    formdata := r.MultipartForm
    if(logActivity){fmt.Println("MULTIPART FILES MAP: ",formdata.File,":",len(formdata.File))}
    
    //ITERATE ALL FILES
    i := 1
    for _, fileHandlers := range formdata.File { 
        for _, fileHandler := range fileHandlers {
            
            //CREATE DESTINATION FILE ON DISK
            dst, creationError := os.Create( fmt.Sprintf("%s/RecievedFile%d.txt", sysSettings.DestinationPath,i) )
            if(creationError != nil){
                fileCreationError(w,r,creationError)
                return
            }
            i++

            //OPEN FILE             
            file, openError := fileHandler.Open()
            if openError != nil {
 			    fileOpenError(w,r,openError)
 			    return
 		    }
            
            //COPY FILE TO DESTINATION ON DISK
            io.Copy(dst, file);
 		    defer file.Close()
 		    
        }

 	}
    
}

