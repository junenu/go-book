package main

import (
	"fmt"
	"time"
)

func main() {

	// timeパッケージで定義されている曜日を日曜日から月曜日までループ
	for day := time.Sunday; day <= time.Saturday; day++ {

		// 平日か休日かを出力
		switch day {
		case time.Sunday:
			fallthrough // 続けて下の「case time.Saturday」が実行される

		case time.Saturday:
			fmt.Println(day, "休日")

		default:
			fmt.Println(day, "たぶん平日")
		}
	}
}
