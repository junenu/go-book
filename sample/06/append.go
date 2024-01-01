package main

import "fmt"

func main() {

	// スライスを宣言
	s1 := []int{1, 2, 3, 4}

	// スライスに要素を追加
	s2 := append(s1, 5, 6)

	// スライスにスライスを追加
	s3 := append(s2, s1...) // スライスの後ろに...が必要

	// 出力
	fmt.Println(s3)

	// appendの特殊な使い方
	var b1 []byte
	b2 := append(b1, "abc"...) // byteスライスに文字列を追加できる
	fmt.Println(b2)
}
