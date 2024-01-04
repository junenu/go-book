package main

import (
	"fmt"
	"time"
)

func main() {
	// 日付の文字列
	dateStr := "2024/4/29"

	// time.Parseを使用してtime.Timeオブジェクトに変換
	date, err := time.Parse("2006/1/2", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// 年、月、日を取得
	year, month, day := date.Year(), date.Month(), date.Day()

	// 出力
	fmt.Printf("%d,%d,%d\n", year, month, day)
}
