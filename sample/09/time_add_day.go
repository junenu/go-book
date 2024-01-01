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

	// 日付を加算(1年と2日後)
	add := t.AddDate(1, 0, 2)

	// 日付を加算(1年と2日前)
	sub := t.AddDate(-1, 0, -2)

	// 計算後の日時を出力
	fmt.Println(add)
	fmt.Println(sub)

}
