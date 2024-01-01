package main

import "fmt"

func main() {

	// 配列を宣言
	values := [...]int{0, 1, 2, 3, 4}

	// 配列をスライス
	s1 := values[1:3]

	// 出力
	fmt.Println(s1)
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	// スライスをスライス
	s2 := s1[1:4] // キャパシティまでスライス可能

	// 出力
	fmt.Println(s2)
	fmt.Println("len:", len(s2))
	fmt.Println("cap:", cap(s2))

	// キャパシティを指定してスライス
	s3 := s1[1:3:4]

	// 出力
	fmt.Println(s3)
	fmt.Println("len:", len(s3))
	fmt.Println("cap:", cap(s3))
}
