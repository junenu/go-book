package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 同時に行う処理数
const goroutines = 3

func main() {

	// チャネルの作成
	c := make(chan int)

	for i := 0; i < goroutines; i++ {

		// ゴルーチン。この関数のパラメータは送信専用チャネル
		go func(s chan<- int) {

			// ダミーの処理として乱数により0～10秒待機
			time.Sleep(
				time.Duration(rand.Int63n(10)) * time.Second)

			fmt.Println("処理完了")

			// 処理完了をチャネルで伝える
			// 送信する値は何でもよい
			s <- 0
		}(c)
	}

	// ゴルーチンの処理完了待ち
	for i := 0; i < goroutines; i++ {
		<-c
	}

	// 完了
	fmt.Println("すべて完了")
}
