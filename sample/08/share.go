package main

import (
	"fmt"
	"os"
)

// ゴルーチンの数
const goroutines = 10

func main() {

	// 値の共用に使用するチャネルの作成
	counter := make(chan int, 1)

	for i := 0; i < goroutines; i++ {

		// ゴルーチン起動
		go func(counter chan int) {

			// チャネルから値を取り出す
			value := <-counter

			// 値を加算
			value++

			// 出力
			fmt.Println("counter:", value)

			// すべてのゴルーチンの処理が完了したら終了
			if value == goroutines {
				os.Exit(0)
			}

			// 更新した値をチャネルに戻す
			counter <- value

		}(counter)
	}

	// チャネルに初期値を送信
	counter <- 0

	// 無限ループ
	// 環境変数「GOMAXPROCS」に2以上を指定する必要あり
	for {
	}
}
