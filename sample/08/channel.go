package main

import "fmt"

func main() {

	// チャネルの作成
	c := make(chan int)

	// ゴルーチン。この関数のパラメータは送信専用チャネル
	go func(s chan<- int) {

		// チャネルへ0～4の値を順番に送信
		for i := 0; i < 5; i++ {
			s <- i
		}

		// チャネルのクローズ
		close(s)
	}(c)

	// 受信ループ
	for {
		// チャネルからの受信
		// クローズされたか知る必要がなければ
		// 2番目の値は受け取る必要はない
		value, ok := <-c

		// チャネルがクローズされるとokにfalseが返される
		if !ok {
			break
		}

		// 受信した値を出力
		fmt.Println(value)
	}
}
