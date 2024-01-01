package main

import "log"

func main() {

	// 標準エラー出力にログを出力
	log.Print("test", 1, 2, 3)
	log.Println("test", 1, 2, 3)
	log.Printf("test %d %d %d", 1, 2, 3)
}
