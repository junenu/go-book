package main

import "fmt"

func main() {

	// スライスを作成
	s := []string{"a", "b", "c"}

	// スライスをそのまま渡すには...を付ける
	test(s...)

	// こうしたときと渡される値は一緒
	test("a", "b", "c")
}

// 可変長パラメータを持つ関数
func test(s ...string) {

	// スライスの長さと値を出力
	fmt.Println(len(s), s)
}
