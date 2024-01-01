package main

import "fmt"

func main() {

	// for文
	for i := 0; i < 5; i++ {

		// 偶数・奇数判定
		if i%2 != 0 {

			// iの値が奇数のときスキップ(for文に戻る)
			continue
		}

		fmt.Println(i)
	}
}
