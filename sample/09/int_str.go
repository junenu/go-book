package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i64 int64 = 1234567890

	// 10進数の文字列に変換し、出力
	// 2番目のパラメータは基数
	fmt.Println(strconv.FormatInt(i64, 10))

	// 16進数の文字列に変換し、出力
	fmt.Println(strconv.FormatInt(i64, 16))

	var i32 int32 = 1234567890

	// int64型でないときは、事前にint64型に変換して呼び出す
	fmt.Println(strconv.FormatInt(int64(i32), 10))
}
