package main

import (
	"fmt"
	"strings"
)

func main() {

	// 分割対象の文字列
	s := "one,two,three,four,five"

	// カンマで分割
	result := strings.Split(s, ",")

	// 分割結果の出力
	fmt.Printf("%q\n", result)
}
