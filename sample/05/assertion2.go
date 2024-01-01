package main

import "fmt"

func main() {

	// 空インタフェースに、string型の値を格納
	var i interface{} = "test"

	// 型アサーションに成功する例
	s1, ok := i.(string)
	fmt.Println(s1, ok)

	// 型アサーションに失敗する例
	// ※string型はdummyメソッドを持たないので変換できない
	s2, ok := i.(interface {
		dummy()
	})
	fmt.Println(s2, ok)
}
