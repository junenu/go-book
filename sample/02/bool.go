package main

import "fmt"

func main() {

	// bool型の変数を宣言
	var b bool

	// bool型の変数にリテラル定数trueとfalseを代入
	b = true
	b = false

	// bool型の変数に論理演算した結果を代入
	b = true || false

	// 出力
	fmt.Println(b)
}
