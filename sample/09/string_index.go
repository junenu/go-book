package main

import (
	"fmt"
	"strings"
)

func main() {

	// 文字列内の「Go」の位置(0～)を取得する
	fmt.Println(strings.Index("The Go Programming Language", "Go"))

	// 見つからないときは-1が返される
	fmt.Println(strings.Index("The Go Programming Language", "xxx"))
}
