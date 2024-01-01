package main

import "fmt"
import "os"

func main() {

	// ファイルのオープン
	file, err := os.Open("test.txt")

	// オープンに失敗したか判定
	if err != nil {

		// エラーの詳細情報を出力
		fmt.Println(err.Error())

		// 終了
		os.Exit(1)
	}

	// ファイルのクローズ
	file.Close()

	fmt.Println("OK")
}
