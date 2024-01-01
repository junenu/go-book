package main

import "os"

func main() {

	// カレントディレクトリ直下に、ディレクトリを作成
	os.Mkdir("newdir", 0777)

	// MkdirAllを使うと途中のディレクトリも作成される
	os.MkdirAll("a/b/c", 0777)
}
