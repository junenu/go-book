package main

import "fmt"

// 四則演算を行う関数の宣言。4つの戻り値を持つ関数
func calc(a int, b int) (int, int, int, float32) {

	// 戻り値を返す(足し算、引き算、掛け算、割り算)
	return a + b, a - b, a * b, float32(a) / float32(b)
}

// main関数はパラメータ、戻り値ともに持たない
func main() {

	// calc関数を呼び出し、多値を受け取る
	add, sub, mul, div := calc(1, 2)

	// 関数から返された戻り値を出力
	fmt.Println(add, sub, mul, div)
}
