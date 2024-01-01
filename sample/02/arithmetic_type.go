package main

import "fmt"

// 型を持たない定数
const C1 = 12345

// 型を持つ(int型)定数
const C2 int = 34567

func main() {

	// C2がint型なので、型を持たないC1もint型として計算される
	var x int = C1 * C2

	// 出力
	fmt.Println("12345 * 34567 =", x)

	// 型が異なる値を計算するときは変換が必要
	var a int32 = 123
	var b int64 = 234

	// 出力
	fmt.Println("123 + 234 =", a+int32(b))
}
