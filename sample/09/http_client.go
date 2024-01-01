package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	// HTTP接続を開始
	response, err := http.Get("http://golang.jp/hello.html")

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// レスポンスのステータスを出力
	fmt.Println("status:", response.Status)

	// レスポンスのボディをすべて読み込む
	body, err := ioutil.ReadAll(response.Body)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ボディを出力
	fmt.Println(string(body))
}
