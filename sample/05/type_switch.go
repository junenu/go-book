package main

import "fmt"

func main() {
	showType(nil)
	showType(12345)
	showType("abcdef")
	showType(3.14)
}

// 型を判定する関数
func showType(x interface{}) {

	// 型switch
	switch x.(type) {

	// nilも使用可能
	case nil:
		fmt.Println("nil")

	// 整数
	case int, int32, int64: // 型を複数指定
		fmt.Println("整数")

	// 文字列
	case string:
		fmt.Println("文字列")

	// その他
	default:
		fmt.Println("不明")
	}
}
