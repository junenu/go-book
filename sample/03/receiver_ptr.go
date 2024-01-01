package main

import "fmt"

// int型からmyType型を宣言
type myType int

// レシーバが値(非ポインタ)のメソッド
func (value myType) setByValue(newValue myType) {

	// ここで値を代入しても無意味
	value = newValue
}

// レシーバがポインタのメソッド
func (value *myType) setByPointer(newValue myType) {

	// 代入した値は反映される
	*value = newValue
}

func main() {

	// myType型の変数を宣言
	var x myType = 0

	// このメソッドのレシーバは値なので、値は変更できない
	x.setByValue(1)
	fmt.Println(x)

	// このメソッドのレシーバはポインタなので、値は変更できる
	x.setByPointer(2)
	fmt.Println(x)
}
