package main

import "fmt"

// 演算インタフェース型
type Calculator interface {

	// 関数の定義
	Calculate(a int, b int) int
}

// 足し算型
type Add struct {
	// フィールドは持たない
}

// Add型にCalculatorインタフェースのCalculate関数を実装
func (x Add) Calculate(a int, b int) int {

	// 足し算
	return a + b
}

// 引き算型
type Sub struct {
	// フィールドは持たない
}

// Sub 型にCalculatorインタフェースのCalculate関数を実装
func (x Sub) Calculate(a int, b int) int {

	// 引き算
	return a - b
}

func main() {

	// Calculatorインタフェースを実装した型の変数を宣言
	var add Add // 足し算
	var sub Sub // 引き算

	// Calculatorインタフェース型の変数を宣言
	var cal Calculator

	// Add型の値を代入
	cal = add

	// インタフェース経由でメソッドを呼び出す
	fmt.Println("和:", cal.Calculate(1, 2))

	// Sub型の値を代入
	cal = sub

	// インタフェース経由でメソッドを呼び出す
	fmt.Println("差:", cal.Calculate(1, 2))
}
