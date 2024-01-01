package main

import "fmt"

func main() {

	// 配列を宣言
	values := [...]int{0, 1, 2, 3, 4}

	// スライスして関数に受け渡す
	double(values[:])

	// 出力。double関数で加えた変更が反映されている
	fmt.Println(values)
}

// スライスの要素の値を倍にする関数
func double(values []int) {

	// 要素を倍にする
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}
