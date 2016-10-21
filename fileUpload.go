package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  "path/filepath"
)

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

func createFileUploadRequest(targetURL string, paramName, path string,fileRepeat int) (*http.Request, error) {
  
  //GET FILE HANDLER
  file, err := os.Open(path)
  if err != nil {
      return nil, err
  }
  defer file.Close()

  //CREATE BODY
  requestBody := &bytes.Buffer{}

  //CREATE WRITER FOR WRITER
  writer := multipart.NewWriter(requestBody)

 i := 0
 for(i<fileRepeat){
    //WRITE PART
    part, err := writer.CreateFormFile(fmt.Sprintf("%s_%d",paramName,i), filepath.Base(path))
    if err != nil {
        return nil, err
    }
    _, err = io.Copy(part, file)
    i++
 }

  //CLOSE WRITER
  err = writer.Close()
  if err != nil {
      return nil, err
  }

  //CREATE REQUEST
  req, err := http.NewRequest("POST", targetURL, requestBody)
  req.Header.Set("Content-Type", writer.FormDataContentType())
  return req, err
}


func main2() {
  //path, _ := os.Getwd()
  //path += "/test.pdf"

  request, err := createFileUploadRequest("http://localhost:8080/assets",  "file", "C://test.txt",500)
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
      fmt.Println(resp.StatusCode)
      fmt.Println(resp.Header)
      fmt.Println(body)
  }
}
