package main

import "fmt"
import "unicode/utf8" // 「utf8」パッケージをインポート

func main() {

	// string型の変数を宣言
	var ja string = "Go言語"

	// 文字数を出力
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja))
}
