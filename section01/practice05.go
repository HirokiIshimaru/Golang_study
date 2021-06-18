package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列型
	fmt.Println("Hello World")

	// 1文字目を取得
	fmt.Println(string("Hello World"[0]))

	// 1文字目を再代入
	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	// worldが含まれているかチェック
	fmt.Println(strings.Contains(s, "World"))

	// 改行
	fmt.Println(`test
	test`)
}
