package main

import "fmt"

func main() {

	// f1の戻り値をf2のパラメータに使用する
	f2(f1())
}

// 多値を返す関数
func f1() (int, string, float32) {
	return 0, "xyz", 3.14
}

// f1の戻り値と同じパラメータを持つ関数
// パラメータcはf1の戻り値(float32)と型が異なるが、interaface{}型には
// どんな値も代入可能なので呼び出せる
// ※interfaceについてはインタフェースの章で説明します
func f2(a int, b string, c interface{}) {
	fmt.Println(a, b, c)
}
