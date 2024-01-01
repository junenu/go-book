package main

import "fmt"

func main() {

	str := "abcあいう"

	// for文(i:インデックス、u:Unicodeコードポイント)
	for i, u := range str {
		fmt.Println(i, u)
	}
}
