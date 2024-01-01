package main

import "fmt"

func main() {
	// Printlnを遅延指定
	defer fmt.Println("defer")

	f1()
}

func f1() {
	// パニック
	panic("パニック発生")
}
