package main

import "fmt"

func main() {

	// 空インタフェースに、string型の値を格納
	var i interface{} = "test"

	// 型アサーションを使いstring型へ
	var s string = i.(string)

	// 出力
	fmt.Println(s)
}
