package main

import (
    "net/http"
    "fmt"
    "io"
    "os"
    "mime/multipart"
    "sync"
    "time"
)

var mutex = &sync.Mutex{}

func handleUpload(w http.ResponseWriter, r *http.Request){
    startTime := time.Now()

	doUpload(w,r)
	w.WriteHeader(http.StatusOK)

     elapsed := time.Since(startTime)
    if(logActivity){fmt.Println("UPLOAD TOOK ", elapsed)}

}


func flushDownloadDirectory(w http.ResponseWriter, r *http.Request){
    sysSettings := GetSettings()
    if(logActivity){fmt.Println("FLUSH REQUESTED: All files in ",sysSettings.DestinationPath," will be deleted...")}
    os.RemoveAll(sysSettings.DestinationPath)
    if(logActivity){fmt.Println("DONE")}
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
    var i uint32
    i = 1  

    for key, fileHandlers := range formdata.File { 
        for _, fileHandler := range fileHandlers {
            
            if(logActivity){fmt.Println("RECIEVED FILE: ",key)}
            go createFileOnDisk(w,r,fileHandler,sysSettings.DestinationPath,key, &i)

        }

 	}
    
}

func createFileOnDisk( w http.ResponseWriter, r *http.Request, fileHandler *multipart.FileHeader, destinationPath , fileName string, iterationIndex *uint32){
  
   mutex.Lock()
     //CREATE DESTINATION FILE ON DISK
        dateString := time.Now().Local().Format("2006-01-02")
        newFile, creationError := os.Create( fmt.Sprint(destinationPath,"/RecievedFile_",dateString,"_",*iterationIndex,"_",fileName) )
        if(creationError != nil){
            fileCreationError(w,r,creationError)
            return
        }
    *iterationIndex++
    mutex.Unlock()

    //OPEN FILE             
        fileToUpload, openError := fileHandler.Open()
        if openError != nil {
            fileOpenError(w,r,openError)
            return
        }
        
    //COPY FILE TO DESTINATION ON DISK
        io.Copy(newFile, fileToUpload);
        defer fileToUpload.Close()
        defer newFile.Close()
}

