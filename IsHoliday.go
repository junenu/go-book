package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	IsHoliday("20200101")
	IsHoliday("20080915")
}

func IsHoliday(object_day string){
	h, err := ioutil.ReadFile(HOLIDAY_JSON)
	if err != nil {
		log.Fatal(err)
	}

	var holidays map[string]HolidayInfo
	err = json.Unmarshal(h, &holidays)
	if err != nil {
		log.Fatal(err)
	}

	for date, info := range holidays {
		if object_day == date {
			fmt.Printf("%d年%d月%d日は%sで祝日です\n", info.Year, info.Month, info.Day, info.Name)
		} else {
			continue
		}
	}
	return false
}
