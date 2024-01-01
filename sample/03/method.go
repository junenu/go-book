package main

import "fmt"

// int型から新たにmyType型を宣言
type myType int

// myType型をレシーバに持つ関数、すなわちmyType型のメソッドを宣言
func (value myType) println() {

	// レシーバの値を出力
	fmt.Println(value)
}

func main() {

	// myType型の変数を宣言
	var z myType = 1234

	// myType型のメソッドを呼び出す
	z.println()
}
