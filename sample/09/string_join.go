package main

import (
	"fmt"
	"strings"
)

func main() {

	// 演算子による結合
	s := "alpha"
	s = s + "beta"
	s += "gamma"

	// 結合結果の出力
	fmt.Println(s)

	// 文字列スライスの用意
	arr := []string{"alpha", "beta", "gamma"}

	// 文字列スライスの結合(カンマで結合)
	fmt.Println(strings.Join(arr, ","))
}
