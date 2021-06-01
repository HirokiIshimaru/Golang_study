package main

import (
	"fmt"
)

// 変数宣言をまとめることができる
// varはmain関数の外で宣言することができる
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	//varなしで変数を宣言する(main関数の外では宣言できない)
	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

// 変数宣言
func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
