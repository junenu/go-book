package main

import (
	"fmt"
	"strings"
)

func main() {

	// 文字列の準備
	s := " xyz "

	// 前後トリム
	fmt.Printf("[%s]\n", strings.TrimSpace(s))

	// 左右どちらかだけトリムする
	fmt.Printf("[%s]\n", strings.TrimLeft(s, " "))
	fmt.Printf("[%s]\n", strings.TrimRight(s, " "))
}
