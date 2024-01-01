package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 空のリストを2つ作成
	l1 := list.New()
	l2 := list.New()

	// リストに要素を追加
	for i := 0; i < 5; i++ {
		l1.PushBack(i)
		l2.PushBack(i)
	}

	// l1にl2の内容を追加
	l1.PushBackList(l2)

	// リストの要素をイテレート
	for e := l1.Front(); e != nil; e = e.Next() {

		// Value関数で要素が得られる
		fmt.Println(e.Value)
	}
}
