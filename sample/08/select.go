package main

import (
	"fmt"
	"os"
)

func main() {

	// チャネルの作成
	ch1 := make(chan int)
	ch2 := make(chan string)

	// ゴルーチン起動
	go func() {

		for i := 0; i < 10; i++ {

			// ch1またはch2への送信
			select {
			// ch1への送信
			case ch1 <- 0:
				fmt.Println("ch1への送信完了")

			// ch2への送信
			case ch2 <- "test":
				fmt.Println("ch2への送信完了")
			}
		}

		// 終了
		os.Exit(0)
	}()

	// 無限ループ
	for {

		// ch1またはch2からの受信
		select {
		// ch1からの受信
		case val := <-ch1:
			fmt.Println("ch1からの受信完了:", val)

		// ch2からの受信
		case text := <-ch2:
			fmt.Println("ch2からの受信完了:", text)

		}
	}
}
