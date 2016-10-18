package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "sync"
    )

// location of the JSON used to populate the config settings 
var settingsLocationPath = "C:/TST/"

//Global singleton variable that holds the settings
var systemSettings *Settings


var once sync.Once


type Settings struct{
    mutex    sync.RWMutex
    DestinationPath string
    ParsingMemoryLimit int64
} 


func (self *Settings) Set(destination string, memoryLimit int64) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.DestinationPath = destination
    self.ParsingMemoryLimit = memoryLimit
}

//Function to load settings from JSON 
func GetSettings() (*Settings){

	//if systemSettings == nil {
    once.Do(func() {

        settings, loadSettingsError := LoadSettings()
        if(loadSettingsError != nil){
            fmt.Println("NO SETTINGS FOUND, USING DEFAULT SETINGS")
            systemSettings.DestinationPath = "C:"
            systemSettings.ParsingMemoryLimit = 200000
        }
        systemSettings = &settings

	})

    return systemSettings
}


func LoadSettings()(Settings,error){

    var settings Settings
    var nilError error
    rawFile, fileLoadError := ioutil.ReadFile(fmt.Sprintf("%s/Settings.json",settingsLocationPath))
    if fileLoadError != nil {
        fmt.Println("PROBLEM LOADING FILE: ",fileLoadError.Error())
        return settings, fileLoadError
    }
    json.Unmarshal(rawFile, &settings)
    if(logActivity){fmt.Println("Loaded Settings :  ",settings)}
    return settings, nilError
}


//Save the recieved settings in JSON
func SaveToSettings(settingToExport Settings){
    
    jason, jsonParsingError := json.Marshal(settingToExport)
    if(jsonParsingError != nil){
        fmt.Println("JSON Parsing Error: ",jsonParsingError.Error())
        return
    }
    fileWritingError := ioutil.WriteFile(fmt.Sprintf("%s/Settings.json",settingsLocationPath), jason, 0644)
    if(fileWritingError != nil){
        fmt.Println("File Writing Error: ",fileWritingError.Error())
        return
    }
}