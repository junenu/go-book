package main

import "fmt"

func main() {

	// ループ
	for i := 0; i < 5; i++ {

		// switch文
		switch i {
		case 0:
			fmt.Println("0")
		case 1, 2: // カンマ区切りで複数の式(値)を記述可能
			fmt.Println("1または2")
		default:
			fmt.Println("その他")
		}
	}
}
