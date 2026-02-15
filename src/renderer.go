package main

import (
	"strconv"
	"strings"
	"fmt"
	"time"
	"github.com/fatih/color"
)

func renderHeader(month string, year int){

	var r,g,b int
	var r2,g2,b2 int
	
	//convert string rgb to int
	_, err := fmt.Sscanf(Colors.headerBox, "%d, %d, %d", &r,&g,&b) 
	if(err != nil){
		fmt.Print(err)
	}

	_, err = fmt.Sscanf(Colors.headerText, "%d,%d,%d", &r2,&g2,&b2)
	if(err != nil){
		fmt.Print(err)
	}

	//define colors as prinf function
	headBorder := color.RGB(r,g,b).PrintfFunc()
	headTitle := color.RGB(r2,g2,b2).PrintfFunc()
	
	title := month + " " + strconv.Itoa(year);
	width := 38; //box width
	leftPadding := (width - len(title)) / 2
	rightPadding := (width - len(title)) - leftPadding


	//print header with defined colors

	headBorder("╭%s╮\n", strings.Repeat("─", width))

	headBorder("│%s", strings.Repeat(" ", leftPadding))
	headTitle(title)
	headBorder("%s│\n", strings.Repeat(" ", rightPadding))

	headBorder("╰%s╯\n", strings.Repeat("─", width))
}

func renderCalendar(currentDay int, currentMonth time.Month, currentYear int){

	var r,g,b int
	var r2,g2,b2 int
	var r3,g3,b3 int
	var r4,g4,b4 int
	_, err := fmt.Sscanf(Colors.passedDays, "%d, %d, %d", &r,&g,&b) 
	if(err != nil){
		fmt.Print(err)
	}
	_, err = fmt.Sscanf(Colors.regularDays, "%d, %d, %d", &r2,&g2,&b2) 
	if(err != nil){
		fmt.Print(err)
	}
	_, err = fmt.Sscanf(Colors.currentDay, "%d, %d, %d", &r3,&g3,&b3) 
	if(err != nil){
		fmt.Print(err)
	}
	_, err = fmt.Sscanf(Colors.weekDaysName, "%d, %d, %d", &r4,&g4,&b4) 
	if(err != nil){
		fmt.Print(err)
	}

	passedDaysColor := color.RGB(r,g,b).PrintfFunc()
	regularDaysColor := color.RGB(r2,g2,b2).PrintfFunc()
	currentDayColor := color.RGB(r3,g3,b3).PrintfFunc()
	weekDaysNameColor := color.RGB(r4,g4,b4).PrintfFunc()

	weekDays := []string{"su", "mo", "tu", "we", "th", "fr", "sa"};

	daysCount := time.Date(currentYear, currentMonth+1, 0, 0,0,0,0, time.Local).Day(); //number of days in current month
	firstWeekday := int(time.Date(currentYear, currentMonth, 1,0,0,0,0, time.Local).Weekday()) //first weekday of the month

	var calendar [6][7]int

	row := 0;
	col := firstWeekday;
	day := 1;

	outer:
	for row < 6 {
		for col < 7{
			if( day <= daysCount){
				calendar[row][col] = day;
				day++;
				col++;

				if(col == 7){
					col = 0;
					row++;
				}
			}else{
				break outer;
			}
		}
	}

	for _, d := range weekDays{
		weekDaysNameColor("%s ", d)
	}

	fmt.Print("\n")

	for i:=0; i<6; i++{
		for j:=0; j<7; j++{
			display := calendar[i][j];
			if(display == 0){
				// fmt.Print("-- ")
				regularDaysColor("-- ")
			}else{
				if(display == currentDay){
					// fmt.Printf("%2d*", display)
					currentDayColor("%2d*", display)
				}else{
					if(display < currentDay){
						passedDaysColor("%2d ", display)
						continue
					}
					// fmt.Printf("%2d ", display)
					regularDaysColor("%2d ", display)
				}
			}
		}
		fmt.Print("\n")
	}
}
