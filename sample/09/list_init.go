package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 空のリストを作成
	l := list.New()

	// リストに要素を追加
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}

	// リストを空にする
	l.Init()

	// リストの要素数を取得する
	fmt.Println("要素数:", l.Len())
}
