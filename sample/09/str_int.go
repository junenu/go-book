package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 文字列をint64に変換	 
	// 2番目のパラメータは基数、10進数なら10を指定
	// 3番目のパラメータはビット数(int:0, int8:8, int16:16, int32:32, int64:64)
	i64, err := strconv.ParseInt("1234567890", 10, 0)

	// int型に変換し、出力
	fmt.Println(int(i64), err)

	// 文字列をint32型に変換
	i64, err = strconv.ParseInt("1234567890", 10, 32)
	fmt.Println(int32(i64), err)
}
