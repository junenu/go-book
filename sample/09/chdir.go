package main

import (
	"fmt"
	"os"
)

func main() {

	// カレントディレクトリを取得し出力
	current, _ := os.Getwd()
	fmt.Println(current)

	// カレントディレクトリを親ディレクトリに変更
	os.Chdir("..")

	// 再びカレントディレクトリを出力
	current, _ = os.Getwd()
	fmt.Println(current)
}
