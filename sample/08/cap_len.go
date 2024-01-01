package main

import "fmt"

func main() {

	// チャネルの作成
	c := make(chan int, 10)

	// 値をひとつ送信
	c <- 0

	// キャパシティと現在の要素数を出力
	fmt.Println("cap:", cap(c))
	fmt.Println("len:", len(c))
}
