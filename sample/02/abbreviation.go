package main

import "fmt"

func main() {

	// 宣言
	a, b := 1, 2

	// cはまだ宣言されていないので:=が使用可能
	a, b, c := 3, 4, 5

	fmt.Println(a, b, c)
}
