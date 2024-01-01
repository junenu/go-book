package main

import (
	"log"
	"os"
)

func main() {

	// ログファイルのオープン(追記モード)
	file, _ := os.OpenFile("test.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// ファイルのクローズ(遅延実行)
	defer func() {
		file.Close()
	}()

	// 標準ロガーの出力先を切り替え
	log.SetOutput(file)
	log.Print("test")

	// 新規にロガーを作成
	// 2番目のパラメータは、ログの各行の先頭に出力するプレフィックス
	// 3番目のパラメータは、ログのフォーマット
	logger := log.New(file, "[mylogger]", log.LstdFlags)
	logger.Print("test")
}
