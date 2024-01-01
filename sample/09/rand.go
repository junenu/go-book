package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 乱数の初期化(現在時刻を乱数の素として使用)
	rand.Seed(time.Now().UnixNano())

	// int型の乱数
	fmt.Println(rand.Int())

	// float32型の乱数
	fmt.Println(rand.Float32())

	// float64型の乱数
	fmt.Println(rand.Float64())
}
