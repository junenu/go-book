package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ゴルーチンの数
const goroutines = 10

// 同時実行できる処理数上限
const maxProcesses = 3

func main() {

	// セマフォ代わりのチャネルの作成(キャパシティに処理上限を設定)
	semaphore := make(chan int, maxProcesses)

	// 完了通知用チャネルの作成
	notify := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {

		// ゴルーチン起動
		go func(no int, semaphore chan int, notify chan<- int) {

			// 自分の番が来るのを待つ
			semaphore <- 0

			// ダミーの処理として乱数により0～3秒待機
			time.Sleep(
				time.Duration(rand.Int63n(3)) * time.Second)

			fmt.Println("処理完了", no)

			// 待機中のゴルーチンに処理を明け渡す
			<-semaphore

			// 処理完了を通知
			notify <- no
		}(i, semaphore, notify)
	}

	// ゴルーチンの処理完了待ち
	for i := 0; i < goroutines; i++ {
		<-notify
	}

	// 完了
	fmt.Println("すべて完了")
}
