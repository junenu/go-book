package main

import "fmt"

func main() {

	// ループ(-2～2)
	for i := -2; i <= 2; i++ {

		// switch文
		switch true {
		case i > 0: // 正の値
			fmt.Println("+")
		case i < 0: // 負の値
			fmt.Println("-")
		default: // その他(ゼロ)
			fmt.Println("0")
		}
	}
}
