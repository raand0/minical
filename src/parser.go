package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(file *os.File){

	//read line by line and assign values to Colors
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := strings.TrimSpace(scanner.Text()) //current line
		if(line == "" || strings.HasPrefix(line, "#")){
			continue
		}

		parts := strings.SplitN(line, "=", 2) //split the line
		if(len(parts) != 2){
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch(key){
			case "headerText": configs.headerText = value;
			case "headerBox": configs.headerBox = value;
			case "currentDay": configs.currentDay = value;
			case "passedDays": configs.passedDays = value;
			case "regularDays": configs.regularDays = value;
			case "weekDaysName": configs.weekDaysName = value;
			case "currentDaySymbol": configs.currentDaySymbol = value;
			case "emptyCellSymbol": configs.emptyCellSymbol = value;
			case "headerBoxWidth": configs.headerBoxWidth, _ = strconv.Atoi(value)
			case "headerBorderX": configs.headerBorderX = value
			case "headerBorderY": configs.headerBorderY = value
			case "headerBorderCLB": configs.headerBorderCLB = value
			case "headerBorderCLU": configs.headerBorderCLU = value
			case "headerBorderCRB": configs.headerBorderCRB = value
			case "headerBorderCRU": configs.headerBorderCRU = value
		}
	}

	err := scanner.Err()
	if(err != nil){
		log.Fatal(err)
	}
}
