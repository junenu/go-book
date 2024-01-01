package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	// 10秒スリープ
	time.Sleep(10 * time.Second)

	fmt.Println(time.Now())

	// 1分スリープ
	time.Sleep(1 * time.Minute)

	fmt.Println(time.Now())
}
