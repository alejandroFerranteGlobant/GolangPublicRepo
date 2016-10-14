package main

import (
    "net/http"
    "fmt"
    "strings"
    "io"
    "os"
)

//-******************
//MAIN FUNCTIONS
//-******************

func handleGet(w http.ResponseWriter, r *http.Request){
    if(LOG_ACTIVITY){
        fmt.Println("Handling GET")
    }
    handlePost(w, r)
}

func handlePost(w http.ResponseWriter, r *http.Request){
    if(LOG_ACTIVITY){
        fmt.Println("Handling POST")
    }

    parameters := strings.Split(strings.Replace(r.URL.String(), " ", "", -1),"/")
    
    //CHECK EMPTY URL
    if(len(parameters) == 0){
        emptyURLError(w)
        return
    }

    switch actionRequested := parameters[1]; actionRequested {
		case "upload":
            if(LOG_ACTIVITY){ fmt.Println("Upload requested") }
			
            //CHECK CORRECT DESTINATION
            
             
            doUpload(w,r)
    }

}

//-******************
//AUXILIAR FUNCTIONS
//-******************

func doUpload(w http.ResponseWriter, r *http.Request){

    //EXRACT FILE FROM HTTP REQUEST

        //recievedFile, header, err := r.FormFile("file")
        recievedFile, exractionError, _ := r.FormFile("file")

        if(recievedFile == nil || exractionError != nil){
            if(LOG_ACTIVITY){fmt.Println("NO FILE RECIEVED")}
            return
        }


        defer recievedFile.Close()
        if(LOG_ACTIVITY){fmt.Println("FILE RECIEVED: ",recievedFile)}


    //CREATE EMPTY FILE
        outputFile, fileCreationError := os.Create("//C:/newFile")  //TODO: Change to use recieved path
        
        if(fileCreationError != nil){
            failedFileCreationError()
        }
        
    //COPY FILE
        _,copyError := io.Copy(outputFile, recievedFile)
        
        if(copyError != nil){
            failedCoyingFileError()
        }

}

