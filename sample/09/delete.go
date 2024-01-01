package main

import "os"

func main() {

	// ディレクトリを作成
	os.MkdirAll("a/b/c/d", 0777)
	os.MkdirAll("x/y/z", 0777)

	// ディレクトリを削除
	// ※実行後「a/b/c」は残る
	os.Remove("a/b/c/d")

	// RemoveAllを使うと指定したディレクトリ以下がすべて削除される
	os.RemoveAll("x")

	// 空のファイルを作成
	file, _ := os.Create("test")
	file.Close()

	// ファイルを削除
	os.Remove("test")
}
