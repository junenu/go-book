package main

import (
	"fmt"
	"math"
)

func main() {

	// 1から5までの平方根を出力
	for i := 1; i <= 5; i++ {
		
		// Sqrtのパラメータはfloat64型
		fmt.Println(i, math.Sqrt(float64(i)))
	}
}
