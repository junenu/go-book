package main

import "fmt"

func main() {

	// fmtを使った出力
	fmt.Print("a", 1, 0.1, "\n")
	fmt.Println("a", 1, 0.1)

	// 組込み関数を使った出力
	print("a", 1, 0.1, "\n")
	println("a", 1, 0.1)
}
