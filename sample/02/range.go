package main

import "fmt"

func main() {

	// 配列の宣言と初期化
	arr := [...]int{0, 1, 2, 3, 4}

	// for文
	for i := range arr {
		fmt.Println(i)
	}
}
