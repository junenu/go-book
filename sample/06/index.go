package main

import "fmt"

func main() {

	// 配列を宣言
	var date [7]string

	// インデックスを使用して要素に値を代入
	date[0] = "日曜日"
	date[1] = "月曜日"
	date[2] = "火曜日"
	date[3] = "水曜日"
	date[4] = "木曜日"
	date[5] = "金曜日"
	date[6] = "土曜日"

	// インデックスを使用して要素から値を取り出す
	for i := 0; i < len(date); i++ {
		fmt.Print(date[i], " ")
	}

	// 出力を見やすくするため改行を出力
	fmt.Println()

	// rangeを使っても要素が取り出せる
	// (rangeから返される一つ目の値はインデックス
	// 使用しないのでブランク識別子に代入)
	for _, value := range date {
		fmt.Print(value, " ")
	}

	// 出力を見やすくするため改行を出力
	fmt.Println()
}
