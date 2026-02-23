package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/gdamore/tcell/v2"
)

var configs struct {
	// Raw strings from config file
	headerText   string
	headerBox    string
	currentDay   string
	passedDays   string
	regularDays  string
	weekDaysName string
	footerColor  string

	// Parsed colors
	headerTextColor    tcell.Color
	headerBoxColor     tcell.Color
	currentDayColor    tcell.Color
	passedDaysColor    tcell.Color
	regularDaysColor   tcell.Color
	weekDaysNameColor  tcell.Color
	footerColorParsed  tcell.Color

	// Other config
	currentDaySymbol string
	emptyCellSymbol  string
	headerBoxWidth   int
	headerBorderX    string
	headerBorderY    string
	headerBorderCLB  string
	headerBorderCLU  string
	headerBorderCRB  string
	headerBorderCRU  string
	showFooter       bool
}

// Layout constants
const (
	calendarWidth = 42
	footerWidth   = 45
	headerY       = 3
	calendarY     = 7
	footerY       = 23
)

var userConfigDir string = getUserConfigDir()
var configDir string = filepath.Join(userConfigDir, "minical")
var configFile string = filepath.Join(configDir, "config")

func getUserConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func createConf() {
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
	}

	defaults := generateText()
	file.Write(defaults)
}

func loadConf() {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parse(file)
	parseColors()
}

func ConfExist() {
	_, err := os.Stat(configFile)
	if err == nil {
		loadConf()
	} else {
		createConf()
		loadConf()
	}
}

func parseColors() {
	configs.headerTextColor = mustParseRGB(configs.headerText)
	configs.headerBoxColor = mustParseRGB(configs.headerBox)
	configs.currentDayColor = mustParseRGB(configs.currentDay)
	configs.passedDaysColor = mustParseRGB(configs.passedDays)
	configs.regularDaysColor = mustParseRGB(configs.regularDays)
	configs.weekDaysNameColor = mustParseRGB(configs.weekDaysName)
	configs.footerColorParsed = mustParseRGB(configs.footerColor)
}

func mustParseRGB(s string) tcell.Color {
	var r, g, b int32
	_, err := fmt.Sscanf(s, "%d, %d, %d", &r, &g, &b)
	if err != nil {
		return tcell.ColorWhite
	}
	return tcell.NewRGBColor(r, g, b)
}

func generateText() []byte {
	return []byte(`
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
showFooter = true
footerColor = 255, 255, 255
	`)
}
