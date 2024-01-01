package main

import "fmt"

func main() {

	// for文
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		// ループの終了判定
		if i == 2 {

			// iの値が2のときfor文から抜ける
			break
		}
	}
}
