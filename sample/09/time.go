package main

import (
	"fmt"
	"time"
)

func main() {

	// ローカルのタイムゾーン情報を取得
	loc, _ := time.LoadLocation("Local")

	// Time型の値を作成(2001/3/3 11:22:33.44)
	time := time.Date(2001, 2, 3, 11, 22, 33, 44, loc)

	// 時刻を出力
	fmt.Println(time)
}
