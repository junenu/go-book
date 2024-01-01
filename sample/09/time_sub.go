package main

import (
	"fmt"
	"time"
)

func main() {

	// 現在日時の取得
	start := time.Now()

	// 15秒待つ
	time.Sleep(time.Second * 15)

	// スリープ後の現在日時の取得
	end := time.Now()

	// 経過時間の算出(Duration型が返る)
	sub := end.Sub(start)

	// 経過時間の出力
	fmt.Println(sub)
	fmt.Println(int(sub/time.Second), "秒")
}
