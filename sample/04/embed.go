package main

import "fmt"

// 埋め込まれる側の構造体
type embedded struct {
	i int
}

// embedded型のメソッド
func (x embedded) doSomething() {
	fmt.Println("test.doSomething()")
}

// 埋め込み先の構造体
type test struct {
	embedded // embedded型の埋め込み
}

func main() {

	var x test

	// embedded型に実装されているメソッドに、test型の値でアクセス
	x.doSomething()

	// embedded型のフィールドに、test型の値でアクセス
	fmt.Println(x.i)
}
