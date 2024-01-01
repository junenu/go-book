package main

import "fmt"

// errorインタフェースを実装したMyError型を宣言
type MyError string

func (e MyError) Error() string {
	return string(e)
}

func main() {

	err := do()

	// エラー判定
	if err == nil {
		fmt.Println("エラーなし")
	} else {
		fmt.Println("エラーあり??", err)
	}
}

// 何か処理をし、エラーがおきてなければnilを返す関数
func do() error {
	var ret *MyError = nil

	// 内部の処理は省略。常に正常終了
	return ret
}
