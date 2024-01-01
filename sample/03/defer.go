package main

import "fmt"

func main() {

	fmt.Println("開始")

	// delay関数を遅延実行
	defer delay()

	fmt.Println("終了")
}

// 遅延指定される関数
// ※deferは呼び出し時に指定するので、この関数自体は
// 遅延実行されることを意識していない
func delay() {
	fmt.Println("delayが呼び出されました")
}
