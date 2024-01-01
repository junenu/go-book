package main

import (
	"log"
	"os"
)

func main() {

	// 存在しないファイルのオープン
	_, err := os.OpenFile("test", os.O_RDONLY, 0)
	if err != nil {
		log.Println("エラー", err)

		// 終了コード1で終了
		os.Exit(1)
	}
}
