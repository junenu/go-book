package main

import "fmt"

func main() {

	// string型の変数を宣言
	var en string = "golang"
	var ja string = "Go言語"

	// 文字列の長さを出力
	fmt.Println(en, " len:", len(en))
	fmt.Println(ja, " len:", len(ja))
}
