package main

import (
	"log"
	"time"
	"github.com/gdamore/tcell/v2"
)

func main() {
	//to do
	//create commands functionality
	//make a gui version
	//create a cool default style

	ConfExist()
	iniTerminal()
}

func iniTerminal() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	viewMonth := today.Month()
	viewYear := today.Year()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Init(); err != nil {
		log.Fatal(err)
	}
	defer s.Fini()

	s.SetStyle(tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset))

	render(s, today, viewMonth, viewYear)

	for {
		ev := s.PollEvent()

		switch e := ev.(type) {
		case *tcell.EventKey:
			prevMonth := viewMonth
			prevYear := viewYear

			switch e.Rune() {
			case 'q', 'Q':
				return
			case '[':
				if viewMonth == time.January {
					viewMonth = time.December
					viewYear--
				} else {
					viewMonth--
				}
			case ']':
				if viewMonth == time.December {
					viewMonth = time.January
					viewYear++
				} else {
					viewMonth++
				}
			case '(':
				viewYear--
			case ')':
				viewYear++
			case 't':
				viewMonth = today.Month()
				viewYear = today.Year()
			}

			// Only redraw if something actually changed
			if viewMonth != prevMonth || viewYear != prevYear {
				s.Clear()
				render(s, today, viewMonth, viewYear)
			}

		case *tcell.EventResize:
			s.Clear()
			s.Sync()
			render(s, today, viewMonth, viewYear)
		}
	}
}

func drawText(s tcell.Screen, x, y int, style tcell.Style, text string) {
	originalX := x
	for _, letter := range text {
		if letter == '\n' {
			y++
			x = originalX
			continue
		}
		s.SetContent(x, y, letter, nil, style)
		x++
	}
}
