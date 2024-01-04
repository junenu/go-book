package main
import (
	"fmt"
	"time"
)

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

func main() {
	IsWeekday("20240104")
	IsWeekday("20231231")
}

func IsWeekday(d string){
	date, _ := time.Parse("20060102", d)
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		fmt.Println("週末です")
	} else {
		fmt.Println("平日です！！！")
	}
}