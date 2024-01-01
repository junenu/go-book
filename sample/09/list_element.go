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

	// リストから4番目(インデックス3)の要素を取り出す
	target := getElement(l, 3)

	// 要素の値を出力
	fmt.Println("変更前:", target.Value)

	// 要素の値を入れ替え
	target.Value = "change"

	// リストの要素をイテレート
	for e := l.Front(); e != nil; e = e.Next() {

		// Value関数で要素が得られる
		fmt.Println(e.Value)
	}
}

// リストの要素(Element)を取り出す関数
func getElement(l *list.List, index int) *list.Element {

	// リストの要素をイテレート
	for e, i := l.Front(), 0; e != nil; e = e.Next() {

		// 指定したインデックスと一致したら、要素を返す
		if i == index {
			return e
		}

		i++
	}

	panic("インデックス不正")
}
