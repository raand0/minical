package main

import (
	"log"
	"time"
	"github.com/gdamore/tcell/v2"
);


func main(){
	//to do: 
	//convert program to TUI using tcell

	ConfExist()

	iniTerminal()

}


func iniTerminal(){

	currentDay := time.Now().Day(); //current Day
	currentMonth := time.Now().Month(); //current Month
	currentYear := time.Now().Year(); //current Year

	s, err := tcell.NewScreen();
	if(err != nil){
		log.Fatal(err)
	}

	err = s.Init();
	if(err != nil){
		log.Fatal(err)
	}
	defer s.Fini()

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(style)

	w,_ := s.Size()
	renderHeader(s, (w/2)-(configs.headerBoxWidth/2),3, currentMonth.String(), currentYear)
	renderCalendar(s, (w/2)-(45/2),7, currentDay, currentMonth, currentYear)
	s.Show()

	for {
		ev := s.PollEvent()

		switch e := ev.(type) {
			case *tcell.EventKey:
			if e.Key() == tcell.KeyEscape || e.Rune() == 'q' {
					return
			}
			case *tcell.EventResize:
				    s.Sync()
		}
	}
	
}

func drawText(s tcell.Screen, x, y int, style tcell.Style, text string){
	originalX := x;

	for _, letter := range text{
		if(letter == '\n'){
			y++;
			x = originalX;
			continue
		}

		s.SetContent(x, y, letter, nil, style)
		x++;
	}
}
