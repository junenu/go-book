package main

import (
	"fmt"
	"net/url"
)

func main() {

	// データ
	data := "abcあいう/:"

	// URLエンコード
	enc := url.QueryEscape(data)

	// エンコード結果を出力
	fmt.Println(enc)

	// エンコード結果をURLデコード
	dec, _ := url.QueryUnescape(enc)

	// デコード結果を出力
	fmt.Println(dec)
}
