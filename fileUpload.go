package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  //"path/filepath"
  "strconv"
  "io/ioutil"
)

/*
// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
  file, err := os.Open(path)
  if err != nil {
      return nil, err
  }
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, err := writer.CreateFormFile(paramName, filepath.Base(path))
  if err != nil {
      return nil, err
  }
  _, err = io.Copy(part, file)

  for key, val := range params {
      _ = writer.WriteField(key, val)
  }
  err = writer.Close()
  if err != nil {
      return nil, err
  }

  req, err := http.NewRequest("POST", uri, body)
  req.Header.Set("Content-Type", writer.FormDataContentType())
  return req, err
}
*/
func createFileUploadRequest(targetURL, path string,fileRepeat int) (*http.Request, error) {

  //CREATE BODY
  requestBody := &bytes.Buffer{}

  //CREATE WRITER FOR WRITER
  writer := multipart.NewWriter(requestBody)

 
    files, _ := ioutil.ReadDir(path)
    for _, fileInfo := range files {
        i := 0
        for(i<fileRepeat){

            //WRITE PART
            name := fileInfo.Name()
            fullPath := fmt.Sprint(path,"/",fileInfo.Name())
            part, formCreationError := writer.CreateFormFile( name , fullPath )
            if formCreationError != nil {
                return nil, formCreationError
            }

            //GET FILE HANDLER
            file, fileOpenError := os.Open(fullPath)
            if fileOpenError != nil {
                return nil, fileOpenError
            }
            defer file.Close()

            //COPY CONTENT
            _, fileCopyError := io.Copy(part, file)
            if(fileCopyError != nil){
                return nil, fileCopyError
            }

        i++
        }


    }
  
  //CLOSE WRITER
  writerClosingError := writer.Close()
  if writerClosingError != nil {
      return nil, writerClosingError
  }

  //CREATE REQUEST
  
  req, requestCreationError := http.NewRequest("POST", targetURL, requestBody)
  req.Header.Set("Content-Type", writer.FormDataContentType())
  return req, requestCreationError
  
}


func main2() {
  //path, _ := os.Getwd()
  //path += "/test.pdf"
originalFilePath := os.Args[1]
repetitions,_ := strconv.Atoi(os.Args[2])

fmt.Println("START LOAD TEST WITH REQUEST WITH ",repetitions," REPETITIONS")
  request, err := createFileUploadRequest("http://localhost:8080/assets/", originalFilePath,repetitions)
  if err != nil {
      log.Fatal(err)
  }
  client := &http.Client{}
  resp, err := client.Do(request)
  if err != nil {
      log.Fatal(err)
  } else {
      body := &bytes.Buffer{}
      _, err := body.ReadFrom(resp.Body)
    if err != nil {
          log.Fatal(err)
      }
    resp.Body.Close()

    fmt.Println("STAUS CODE:",resp.StatusCode)
    fmt.Println("RESPONSE HEADER: ",resp.Header)
    fmt.Println("BODY: ",body)
  }
}
