package main

import "fmt"

func main() {

	// 長さ5の配列。各要素はゼロ値で初期化される
	array1 := [5]float32{}

	// 長さ6の配列。要素の不足分はゼロ値で初期化される
	array2 := [6]int{1, 2, 3, 4}

	// 長さに「...」と記述すると要素数が長さとして使用される
	array3 := [...]string{"One", "Two", "Three"}

	// 出力
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}
