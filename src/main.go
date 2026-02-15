package main

import (
	"time"
);


func main(){
	//to do: 
	// create a parser
	// if config file exists load
	// if not create

	ConfExist()

    	currentDay := time.Now().Day(); //current Day
	currentMonth := time.Now().Month(); //current Month
	currentYear := time.Now().Year(); //current Year

	renderHeader(currentMonth.String(), currentYear)
	renderCalendar(currentDay, currentMonth, currentYear)

}

