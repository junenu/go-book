package main

import "fmt"

func main() {

	// ループ用のカウンタ
	i := 0

	// for文
	for {
		fmt.Println(i)
		i++

		// ループの終了判定
		if i == 5 {

			// for文から抜ける
			break
		}
	}
}
