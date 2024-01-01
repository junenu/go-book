package main

import "fmt"

func main() {

LBL:
	// 1番目のループ
	for i := 0; i < 3; i++ {
		fmt.Println(i)

		// 2番目のループ
		for j := 0; j < 3; j++ {
			fmt.Println("   ", j)

			// 条件によりcontinue
			if i == 1 && j == 1 {

				// 外側(1番目)のforへcontinue
				continue LBL
			}
		}
	}
}
