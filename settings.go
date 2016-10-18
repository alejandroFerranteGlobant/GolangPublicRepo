package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    )

var settingsLocationPath = "C:/TST/"

type Settings struct{
    DestinationPath string
    ParsingMemoryLimit int64
} 

func LoadSettings()(Settings,error){

    var settings Settings
    var nilError error
    rawFile, fileLoadError := ioutil.ReadFile(fmt.Sprintf("%s/Settings.json",settingsLocationPath))
    if fileLoadError != nil {
        fmt.Println("PROBLEM LOADIN FILE: ",fileLoadError.Error())
        return settings, fileLoadError
    }
    json.Unmarshal(rawFile, &settings)
    if(logActivity){fmt.Println("Loaded Settings :  ",settings)}
    return settings, nilError
}





//-*****************
//SAMPLE CODE
//-*****************

func SaveXample(){
    Seting := Settings{"C:/TST/", 2000000}
    jason, jsonParsingError := json.Marshal(Seting)
    if(jsonParsingError != nil){
        fmt.Println("JSON Parsing Error")
        return
    }
    fileWritingError := ioutil.WriteFile("C:/TST/newJson.json", jason, 0644)
    if(fileWritingError != nil){
        fmt.Println("File Writing Error")
        return
    }
}