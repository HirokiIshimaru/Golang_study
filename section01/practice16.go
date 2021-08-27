package main

import "fmt"

func main() {
	// ポインタ
	// (1) int型変数ageを定義
	var age int = 30

	// (2) int型のポインタ変数を定義
	var p *int

	// (3) &(アドレス演算子)を使って、nのアドレスを代入
	p = &age

	// (4) pの値(アドレス)を出力する
	fmt.Println("pのアドレス:", p)

	// (5) *を使ってポインタのリテラル（中身）を出力
	fmt.Println("pの中身:", p)

	// (6) 変数ageのリテラル（中身）を出力
	fmt.Println("ageの値:", age)

	// (7) ＆を使うことで変数ageのアドレスを取得できる
	fmt.Println("ageのアドレス:", &age)
}
