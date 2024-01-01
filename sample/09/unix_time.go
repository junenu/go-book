package main

import (
	"fmt"
	"time"
)

func main() {

	// 現在日時の取得
	t := time.Now()

	// UNIX時間に変換
	unix := t.Unix()

	// UNIX時間の出力
	fmt.Println(unix)
}
