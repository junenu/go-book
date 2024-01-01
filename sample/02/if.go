package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {

		// if文(奇数・偶数の判定)
		if i%2 == 0 {
			fmt.Println(i, "は偶数")
		} else {
			fmt.Println(i, "は奇数")
		}
	}
}
