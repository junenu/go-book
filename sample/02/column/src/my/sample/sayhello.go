package sample // 「my/sample」だが最下位のパッケージ名だけ記述する

import "fmt"

func SayHello(who string) {

	// 出力
	fmt.Printf("Hello, %s!\n", who)
}
