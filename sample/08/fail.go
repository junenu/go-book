package main

import (
	"fmt"
	"os"
)

func main() {

	// 関数リテラルを使いゴルーチンを起動
	go func() {
		fmt.Println("Goroutine")

		// プログラム終了
		os.Exit(0)
	}()

	// 無限ループ
	for {
	}
}
