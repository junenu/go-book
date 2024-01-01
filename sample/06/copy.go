package main

import "fmt"

func main() {

	// スライスを作成
	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	// 要素のコピー
	count := copy(dst[2:], src)

	// 出力
	fmt.Println(dst)
	fmt.Println("count:", count)
}
