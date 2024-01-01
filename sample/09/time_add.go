package main

import (
	"fmt"
	"time"
)

func main() {

	// 現在日時の取得
	t := time.Now()

	// 現在日時の出力
	fmt.Println(t)

	// 日時を加算(24:30:05後)
	add := t.Add(time.Hour*24 + time.Minute*30 + time.Second*5)

	// 日時を加算(24:30:05前)
	sub := t.Add(-time.Hour*24 - time.Minute*30 - time.Second*5)

	// 計算後の日時を出力
	fmt.Println(add)
	fmt.Println(sub)

}
