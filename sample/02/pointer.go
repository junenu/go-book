package main

import "fmt"

func main() {

	// int型のポインタ変数を宣言
	var ptr *int

	// int型の変数を宣言
	var i int = 12345

	// アドレス演算子を使いiのアドレスをptrに代入
	ptr = &i

	// 出力
	fmt.Println("iのアドレス:", &i)
	fmt.Println("ptrの値(変数iのアドレス):", ptr)

	fmt.Println("iの値:", i)
	fmt.Println("ポインタ経由のiの値:", *ptr)

	// ポインタ経由でiの値を変更
	*ptr = 999

	// 出力
	fmt.Println("ポインタ経由で変更したiの値:", i)
}
