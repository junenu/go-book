package main

import "fmt"

func main() {

	// AND条件
	fmt.Println("true  && true  =", true && true)   // 真 AND 真
	fmt.Println("true  && false =", true && false)  // 真 AND 偽
	fmt.Println("false && true  =", false && true)  // 偽 AND 真
	fmt.Println("false && false =", false && false) // 偽 AND 偽

	// OR条件
	fmt.Println("true  || true  =", true || true)   // 真 OR 真
	fmt.Println("true  || false =", true || false)  // 真 OR 偽
	fmt.Println("false || true  =", false || true)  // 偽 OR 真
	fmt.Println("false || false =", false || false) // 偽 OR 偽

	// 否定
	fmt.Println("!true  =", !true)  // 真の否定
	fmt.Println("!false =", !false) // 偽の否定
}
