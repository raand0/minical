package main

import (
	"bufio"
	"log"
	"os"
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
			case "headerText": Colors.headerText = value;
			case "headerBox": Colors.headerBox = value;
			case "currentDay": Colors.currentDay = value;
			case "passedDays": Colors.passedDays = value;
			case "regularDays": Colors.regularDays = value;
			case "weekDaysName": Colors.weekDaysName = value;
		}
	}

	err := scanner.Err()
	if(err != nil){
		log.Fatal(err)
	}
}
