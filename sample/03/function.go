package main

import "fmt"

// 足し算を行う関数の宣言
func plus(a int, b int) int {

	// 戻り値を返す
	return a + b
}

// main関数はパラメータ、戻り値ともに持たない
func main() {

	// plus関数を呼び出す
	answer := plus(1, 2)

	// 関数から返された戻り値を出力
	fmt.Println(answer)
}
