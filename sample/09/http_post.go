package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {

	// HTTP接続を開始
	response, err := http.PostForm("http://golang.jp/hello.html",
		url.Values{
			"key1": {"Value1", "Value2"},
			"key2": {"Value"}})

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
