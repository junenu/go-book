package main

import (
	"fmt"
	"strconv"
)

func main() {

	var f64 float64 = 1.23456789

	// 文字列に変換し、出力
	// 2番目のパラメータはフォーマット
	// 3番目のパラメータは精度(e,E:指数あり、f:指数なし)
	// 4番目のパラメータはビット数(32または64)
	fmt.Println(strconv.FormatFloat(f64, 'e', 8, 64))
	fmt.Println(strconv.FormatFloat(f64, 'f', 4, 64))

	var f32 float32 = 1.23

	// float32型のときは、事前にflot64型に変換する
	fmt.Println(strconv.FormatFloat(float64(f32), 'e', 2, 32))
}
