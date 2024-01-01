package main

import "fmt"

// main関数
func main() {

	// 構造体型の変数を宣言
	var x struct {
		i1, i2 int     // int型フィールド
		f      float32 // float32型フィールド
		s      string  // string型フィールド
	}

	// 構造体の各フィールドに値をセット
	x.i1 = 1
	x.i2 = 2
	x.f = 3.14
	x.s = "go"

	// 出力
	fmt.Println(x)
}
