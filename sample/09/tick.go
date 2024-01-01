package main

import (
	"fmt"
	"time"
)

func main() {

	// 5秒ごとに通知されるチャネルを作成
	ch := time.Tick(time.Second * 5)

	// 10回ループ
	for i := 0; i < 10; i++ {

		// チャネルから受信待ち
		t := <-ch

		// ここに一定時間ごとに行う処理を記述
		// 受信した日時を出力
		fmt.Println(t)
	}
}
