package main

import (
	"fmt"
	"math"
)

func main() {

	// 算出する値として円周率を使用する
	val := math.Pi

	fmt.Println("元の値:", val)

	// 切り上げを行う
	fmt.Println("切り上げ:", math.Ceil(val))

	// 切り捨てを行う
	fmt.Println("切り捨て:", math.Floor(val))

	// 値をマイナス値にする
	val *= -1

	fmt.Println("元の値:", val)

	// 切り上げを行う
	fmt.Println("切り上げ:", math.Ceil(val))

	// 切り捨てを行う
	fmt.Println("切り捨て:", math.Floor(val))
}
