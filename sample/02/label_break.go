package main

import "fmt"

func main() {

LBL:
	// ループ
	for i := 0; i < 5; i++ {

		// switch文
		switch {
		case i == 3:
			// forへbreak
			break LBL

		default:
			// 出力
			fmt.Println(i)
		}
	}
}
