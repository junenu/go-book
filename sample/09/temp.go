package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// テンポラリディレクトリ名を取得
	dir := os.TempDir()

	// テンポラリディレクトリ内にテンポラリファイルを作成
	file, _ := ioutil.TempFile(dir, "test")

	// 作成したテンポラリファイル名の出力
	fmt.Println(file.Name())
}
