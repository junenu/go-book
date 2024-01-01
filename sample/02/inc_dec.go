package main

import "fmt"

func main() {

	// 初期値0の変数を2つ宣言
	var inc int = 0
	var dec int = 0

	inc++ // インクリメント
	dec-- // デクリメント

	// 結果を出力
	fmt.Println("0++ =", inc)
	fmt.Println("0-- =", dec)
}
