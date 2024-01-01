package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func main() {

	// 一覧取得先のディレクトリとして$GOROOTを使用する
	goroot := runtime.GOROOT()

	// ディレクトリ内のファイル、ディレクトリ一覧を取得
	fileinfos, _ := ioutil.ReadDir(goroot)

	// 取得した一覧を出力
	for _, fileinfo := range fileinfos {

		// 今回はディレクトリは除外
		if !fileinfo.IsDir() {
			fmt.Println(fileinfo.Name())
		}
	}
}
