package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
)

const (
	HOLIDAY_JSON = "./holidays.json"
)

type HolidayInfo struct{
	Year   int
	Month  int
	Day    int
	Name   string
} 

func main() {
	WhatHoliday("20200101")
}

func WhatHoliday(object_day string){
	h, err := os.ReadFile(HOLIDAY_JSON)
	if err != nil {
		log.Fatal(err)
	}

	var holidays map[string]HolidayInfo
	err = json.Unmarshal(h, &holidays)
	if err != nil {
		log.Fatal(err)
	}

	for date, holiday := range holidays {
		if object_day == date {
			fmt.Println(holiday)
		} else {
			continue
		}
	}
	fmt.Println("No")
}