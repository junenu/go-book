package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// 「go help」を実行するためのCmdを作成
	cmd := exec.Command("go", "help")

	// 実行し、コマンドが出力した結果を取得
	result, err := cmd.Output()

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 実行結果を出力
	fmt.Printf("%s\n", result)
}
