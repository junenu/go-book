package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 文字列をfloat64に変換
	// 2番目のパラメータはビット数(float32:32, float64:64)
	f64, err := strconv.ParseFloat("123.456789", 32)

	// float32型に変換し、出力
	fmt.Println(float32(f64), err)
}
