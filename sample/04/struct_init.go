package main

import "fmt"

// 構造体型の宣言
type Person struct {
	name string
	age  int
}

func main() {

	// 構造体リテラルを使用せず、フィールドを個別に初期化
	var p1 Person
	p1.name = "Jhon"
	p1.age = 23

	// 構造体リテラルで初期化(フィールド名と値のペアを記述)
	p2 := Person{age: 31, name: "Tom"}

	// 構造体リテラルで初期化(フィールドの宣言順に値を記述)
	p3 := Person{"Jane", 42}

	// ポインタも構造体リテラルで作成可能
	p4 := &Person{"Mike", 36}

	// 出力
	fmt.Println(p1, p2, p3, p4)
}
