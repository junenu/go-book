package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// ホスト名をIPアドレスに変換
	addrs, err := net.LookupHost("www.google.com")

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 取得したIPアドレスを出力
	fmt.Printf("%q\n", addrs)
}
