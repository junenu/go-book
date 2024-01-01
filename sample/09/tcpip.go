package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	// ホスト(HTTPサーバ)に接続
	conn, err := net.Dial("tcp", "golang.jp:80")

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 通信のクローズ(遅延実行)
	defer func() {
		conn.Close()
	}()

	// HTTPリクエストの送信
	fmt.Fprintf(conn, "GET /hello.html HTTP/1.1\r\n")
	fmt.Fprintf(conn, "HOST: golang.jp\r\n")
	fmt.Fprintf(conn, "\r\n")

	// HTTPレスポンスの受信
	response, err := ioutil.ReadAll(conn)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// レスポンスをバイトスライスから文字列に変換し出力
	fmt.Println(string(response))
}
