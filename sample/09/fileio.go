package main

import (
	"fmt"
	"os"
)

func main() {
	write()
	read()
}

// 書き込み
func write() {

	// ファイルのオープン
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ファイルのクローズ(遅延実行)
	defer func() {
		file.Close()
	}()

	// 文字列をファイルに書き込む
	file.WriteString("test\n")

	// Fprintfを使っても書き込める
	fmt.Fprintf(file, "test\n")
}

// 読み込み
func read() {

	// ファイルのオープン
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ファイルのクローズ(遅延実行)
	defer func() {
		file.Close()
	}()

	// ファイルから文字列を読み込む
	var str string
	fmt.Fscanf(file, "%s", &str) // 改行またはスペースまで読み込む

	// 読み込んだ文字列の出力
	fmt.Println(str)

	// 以降のデータをbyteスライスに読み込む
	buf := make([]byte, 16)
	size, _ := file.Read(buf)

	// 読み込んだバイナリデータの出力
	fmt.Println(size, buf)
}
