package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main:開始")

	fmt.Println("testを通常の関数として起動")
	test()

	fmt.Println("testをゴルーチンとして起動")
	go test()

	// 3秒待つ
	time.Sleep(3 * time.Second)

	fmt.Println("main:終了")
}

// ゴルーチンとして起動する関数
func test() {

	for i := 0; i < 5; i++ {

		// 連番を出力
		fmt.Println(i)

		// 1秒待つ
		time.Sleep(1 * time.Second)
	}
}
