package main

import (
	"fmt"
	"os"
)

func main() {

	// ディレクトリを作成
	os.MkdirAll("testdir", 0777)

	// ディレクトリの名前変更
	err := os.Rename("testdir", "newdir")

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 空のファイルを作成
	file, _ := os.Create("testfile")
	file.Close()

	// ファイルの名前変更
	err = os.Rename("testfile", "newfile")

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
