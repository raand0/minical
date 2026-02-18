package main

import (
	"log"
	"os"
	"path/filepath"
)

var configs struct{
	headerText string
	headerBox string
	currentDay string
	passedDays string
	regularDays string
	weekDaysName string
	currentDaySymbol string
	emptyCellSymbol string
	headerBoxWidth int
	headerBorderX string
	headerBorderY string
	headerBorderCLB string
	headerBorderCLU string
	headerBorderCRB string
	headerBorderCRU string
}

var userConfigDir string = getUserConfigDir()
var configDir string = filepath.Join(userConfigDir, "minical") //home/.config/minical  %Appdata%/minical
var configFile string = filepath.Join(configDir, "config")

func getUserConfigDir() string{
	dir, err := os.UserConfigDir()
	if(err != nil){
		log.Fatal(err)
	}

	return dir;
}

func createConf(){
	
	err := os.MkdirAll(configDir, 0755)
	if(err != nil){
		log.Fatal(err)
	}

	file, err := os.Create(configFile) //config file
	if(err != nil){
		log.Fatal(err)
	}
	
	//write defaults to file
	defaults := generateText()
	file.Write(defaults)
}

func loadConf(){

	file, err := os.Open(configFile) //get the config file
	if(err != nil){
		log.Fatal(err)
	}
	defer file.Close()

	parse(file)
}

func ConfExist(){
	_, err := os.Stat(configFile)
	if(err == nil){
		loadConf()
	}else {
		createConf()
		loadConf()
	}
}

func generateText() []byte{
	text := []byte(`
headerText = 66, 135, 245
headerBox = 163, 199, 255
currentDay = 112, 255, 200
passedDays = 140, 140, 140
regularDays = 255, 255, 255
weekDaysName = 104, 66, 179
currentDaySymbol = ◆
emptyCellSymbol = ·
headerBoxWidth = 38
headerBorderX = ─
headerBorderY = │
headerBorderCLB = ╰
headerBorderCLU = ╭
headerBorderCRB = ╯
headerBorderCRU = ╮
	`);

	return text;
}
