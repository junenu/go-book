package main

import (
	"fmt"
	"strings"
)

func main() {

	// 文字列の準備
	s := "This is a pen."

	// 大文字に変換
	fmt.Println(strings.ToUpper(s))

	// 小文字に変換
	fmt.Println(strings.ToLower(s))
}
