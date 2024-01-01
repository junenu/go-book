package main

import "fmt"

func main() {

	// 配列を宣言
	x := [5]string{"a", "b", "c", "d", "e"}

	// スライス型の変数を宣言
	var s1 []string

	// 配列全体をスライス
	s1 = x[:]
	fmt.Println(s1)

	// インデックス1～3までをスライス
	s2 := x[1:4]
	fmt.Println(s2)

	// インデックス3～をスライス
	s3 := x[3:]
	fmt.Println(s3)

	// インデックス～3をスライス
	s4 := x[:4]
	fmt.Println(s4)
}
