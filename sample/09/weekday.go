package main

import (
	"fmt"
	"time"
)

func main() {

	// 現在日時の取得
	t := time.Now()

	// 曜日の取得と出力
	switch t.Weekday() {
	case time.Sunday:
		fmt.Println("日曜日")

	case time.Monday:
		fmt.Println("月曜日")

	case time.Tuesday:
		fmt.Println("火曜日")

	case time.Wednesday:
		fmt.Println("水曜日")

	case time.Thursday:
		fmt.Println("木曜日")

	case time.Friday:
		fmt.Println("金曜日")

	case time.Saturday:
		fmt.Println("土曜日")
	}
}
