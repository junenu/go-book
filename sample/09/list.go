package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 空のリストを作成
	l := list.New()

	// リストに要素を追加
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack(3)

	// リストの要素をイテレート
	for e := l.Front(); e != nil; e = e.Next() {

		// Value関数で要素が得られる
		fmt.Println(e.Value)
	}
}
