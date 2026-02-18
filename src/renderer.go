package main

import (
	"strings"
	"strconv"
	"fmt"
	"time"
	"github.com/gdamore/tcell/v2"
)

func renderHeader(s tcell.Screen, x,y int, month string, year int) {

	var r,g,b int32
	var r2,g2,b2 int32
	
	//convert string rgb to int
	_, err := fmt.Sscanf(configs.headerBox, "%d, %d, %d", &r,&g,&b) 
	if(err != nil){
		fmt.Print(err)
	}

	_, err = fmt.Sscanf(configs.headerText, "%d,%d,%d", &r2,&g2,&b2)
	if(err != nil){
		fmt.Print(err)
	}
	
	title := month + " " + strconv.Itoa(year)

	width := configs.headerBoxWidth
	titleLen := len(title)

	leftPadding := (width - titleLen) / 2
	rightPadding := (width - titleLen) - leftPadding

	borderStyle := tcell.StyleDefault.Foreground(
		tcell.NewRGBColor(r, g, b),
	)

	titleStyle := tcell.StyleDefault.Foreground(
		tcell.NewRGBColor(r2, g2, b2),
	)

	drawText(s, x, y, borderStyle,
		fmt.Sprintf("%s%s%s", configs.headerBorderCLU, strings.Repeat(configs.headerBorderX, width), configs.headerBorderCRU))

	drawText(s, x, y+1, borderStyle,
		fmt.Sprintf("%s%s", configs.headerBorderY, strings.Repeat(" ", leftPadding)))

	drawText(s, x+1+leftPadding, y+1, titleStyle, title)

	drawText(s, x+1+leftPadding+titleLen, y+1, borderStyle,
		fmt.Sprintf("%s%s", strings.Repeat(" ", rightPadding), configs.headerBorderY))

	drawText(s, x, y+2, borderStyle,
		fmt.Sprintf("%s%s%s", configs.headerBorderCLB, strings.Repeat(configs.headerBorderX, width), configs.headerBorderCRB))

}

func renderCalendar(s tcell.Screen, x,y int, currentDay int, currentMonth time.Month, currentYear int){

	var r,g,b int32
	var r2,g2,b2 int32
	var r3,g3,b3 int32
	var r4,g4,b4 int32

	_, err := fmt.Sscanf(configs.passedDays, "%d, %d, %d", &r,&g,&b) 
	if(err != nil){
		fmt.Print(err)
	}
	_, err = fmt.Sscanf(configs.regularDays, "%d, %d, %d", &r2,&g2,&b2) 
	if(err != nil){
		fmt.Print(err)
	}
	_, err = fmt.Sscanf(configs.currentDay, "%d, %d, %d", &r3,&g3,&b3) 
	if(err != nil){
		fmt.Print(err)
	}
	_, err = fmt.Sscanf(configs.weekDaysName, "%d, %d, %d", &r4,&g4,&b4) 
	if(err != nil){
		fmt.Print(err)
	}

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

	passedDaysStyle := tcell.StyleDefault.Foreground(tcell.NewRGBColor(r,g,b))
	regularDaysStyle := tcell.StyleDefault.Foreground(tcell.NewRGBColor(r2,g2,b2))
	currentDayStyle := tcell.StyleDefault.Foreground(tcell.NewRGBColor(r3,g3,b3))
	weekDaysStyle := tcell.StyleDefault.Foreground(tcell.NewRGBColor(r4,g4,b4))
	originalX := x
	for _,d := range weekDays{
		drawText(s, x,y, weekDaysStyle, d+"     ")
		drawText(s, x,y+1, weekDaysStyle, "──   ")
		x += 7
	}
	y += 2;
	x = originalX

	for i:=0; i<6; i++{
		for j:=0; j<7; j++{
			display := calendar[i][j];
			if(display == 0){
				drawText(s, x,y, regularDaysStyle, configs.emptyCellSymbol + "      ")
				x += 7;
			}else{
				if(display == currentDay){
					drawText(s, x,y, currentDayStyle, fmt.Sprintf("%2d%s    ", display, configs.currentDaySymbol))
					x += 7;
				}else{
					if(display < currentDay){
						drawText(s, x,y, passedDaysStyle, fmt.Sprintf("%2d     ", display))
						x += 7
						continue
					}
					drawText(s, x,y, regularDaysStyle, fmt.Sprintf("%2d     ", display))
					x += 7;
				}
			}
		}
		y += 2
		x = originalX
	}
}
