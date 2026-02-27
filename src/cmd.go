package main

import (
	"flag"
	"fmt"
	"time"
	"os"
)

func flags(){

	today := flag.Bool("today", false, "Returns current date in YYYY-MM-DD format")
	genConfig := flag.Bool("gen-config", false, "Generates default config file")

	flag.Parse()

	if(*today){
		fmt.Println(currentDate())
		os.Exit(0)
	}
	if(*genConfig){
		fmt.Println(createConf())
		os.Exit(0)
	}
}

func currentDate() string{

	now := time.Now()
	return now.Format("2006-01-02")
}
