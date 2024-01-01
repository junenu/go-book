package main

import "fmt"

func main() {

	// 文字列のスライス
	var x string = "abcde"[1:4]

	// 平仮名はUTF-8で3バイトなので「いう」が取り出せる
	var y string = "あいうえお"[3:9]

	// UTF-8の文字の境界として不適切な値を指定すると文字化けする
	var z string = "あいうえお"[1:4]

	// 出力
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
