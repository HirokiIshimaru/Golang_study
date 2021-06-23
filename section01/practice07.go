package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 型変換

	var x int = 1
	xx := float64(x)                     //ここで型変換
	fmt.Printf("%T %v %f\n", xx, xx, xx) //タイプとvalueとfloat表記を表示

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s) //文字列の型変換
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
