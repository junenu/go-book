package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	// データ
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	// BASE64エンコード
	enc := base64.StdEncoding.EncodeToString(data)

	// エンコード結果を出力
	fmt.Println(enc)

	// エンコード結果をBASE64デコード
	dec, _ := base64.StdEncoding.DecodeString(enc)

	// デコード結果を出力
	fmt.Printf("% x\n", dec)
}
