package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
)

var weekDays [7]string = [7]string{"su", "mo", "tu", "we", "th", "fr", "sa"}

var calendarCache struct {
	month  time.Month
	year   int
	grid   [6][7]int
	valid  bool
}

func getCachedCalendar(month time.Month, year int) [6][7]int {
	if calendarCache.valid && calendarCache.month == month && calendarCache.year == year {
		return calendarCache.grid
	}
	calendarCache.grid = populateCalender(month, year)
	calendarCache.month = month
	calendarCache.year = year
	calendarCache.valid = true
	return calendarCache.grid
}

func render(s tcell.Screen, today time.Time, viewMonth time.Month, viewYear int) {
	w, _ := s.Size()
	xCal := (w / 2) - (calendarWidth / 2)
	xHeader := (w / 2) - (configs.headerBoxWidth / 2)
	xFooter := (w / 2) - (footerWidth / 2)

	renderHeader(s, xHeader, headerY, viewMonth.String(), viewYear)
	renderCalendar(s, xCal, calendarY, today, viewMonth, viewYear)
	renderFooter(s, xFooter, footerY)
	s.Show()
}

func renderHeader(s tcell.Screen, x, y int, month string, year int) {
	borderStyle := tcell.StyleDefault.Foreground(configs.headerBoxColor)
	titleStyle := tcell.StyleDefault.Foreground(configs.headerTextColor).Bold(true)

	title := fmt.Sprintf("%s %d", month, year)
	width := configs.headerBoxWidth
	titleLen := len(title)
	leftPadding := (width - titleLen) / 2
	rightPadding := (width - titleLen) - leftPadding

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

func renderFooter(s tcell.Screen, x, y int) {
	if !configs.showFooter {
		return
	}
	footerStyle := tcell.StyleDefault.Foreground(configs.footerColorParsed)
	drawText(s, x, y, footerStyle, "[] • month   () • year   t • today   q • quit")
}

func populateCalender(currentMonth time.Month, currentYear int) [6][7]int {
	daysCount := time.Date(currentYear, currentMonth+1, 0, 0, 0, 0, 0, time.Local).Day()
	firstWeekday := int(time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local).Weekday())

	var result [6][7]int
	row, col, day := 0, firstWeekday, 1

outer:
	for row < 6 {
		for col < 7 {
			if day <= daysCount {
				result[row][col] = day
				day++
				col++
				if col == 7 {
					col = 0
					row++
				}
			} else {
				break outer
			}
		}
	}
	return result
}

func renderCalendar(s tcell.Screen, x, y int, today time.Time, viewMonth time.Month, viewYear int) {
	weekDaysStyle := tcell.StyleDefault.Foreground(configs.weekDaysNameColor).Dim(true)

	highlightDay := 0
	if viewYear == today.Year() && viewMonth == today.Month() {
		highlightDay = today.Day()
	}

	currentDayStyle := tcell.StyleDefault.Foreground(configs.currentDayColor).Bold(true)
	passedDaysStyle := tcell.StyleDefault.Foreground(configs.passedDaysColor)
	regularDaysStyle := tcell.StyleDefault.Foreground(configs.regularDaysColor)

	calendar := getCachedCalendar(viewMonth, viewYear)

	originalX := x
	for _, d := range weekDays {
		drawText(s, x, y, weekDaysStyle, d+"     ")
		drawText(s, x, y+1, weekDaysStyle, "──   ")
		x += 7
	}
	y += 2
	x = originalX

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			d := calendar[i][j]
			if d == 0 {
				drawText(s, x, y, regularDaysStyle, configs.emptyCellSymbol+"      ")
			} else if d == highlightDay {
				drawText(s, x, y, currentDayStyle, fmt.Sprintf("%2d%s    ", d, configs.currentDaySymbol))
			} else if highlightDay > 0 && d < highlightDay {
				// Only colour as "passed" when we're viewing the current month
				drawText(s, x, y, passedDaysStyle, fmt.Sprintf("%2d     ", d))
			} else {
				drawText(s, x, y, regularDaysStyle, fmt.Sprintf("%2d     ", d))
			}
			x += 7
		}
		y += 2
		x = originalX
	}
}
