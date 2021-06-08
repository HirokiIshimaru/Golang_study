package main

import (
	"fmt"
)

// 数値型（Numeric types）
func main() {
	// var (
	// 	u8  uint8     = 255
	// 	i8  int8      = 127
	// 	f32 float32   = 0.2
	// 	c64 complex64 = -5 + 12i
	// )
	// fmt.Println(u8, i8, f32, c64)

	// fmt.Printf("type=%T value=%v", u8, u8)
	// %vはデフォルトの値を表示、%Tはtypeを表示

	fmt.Println("1+1=", 1+1)
	fmt.Println("10-1=", 10-1)
	fmt.Println("10/2=", 10/2)
	fmt.Println("10/3=", 10/3)
	fmt.Println("10.0/3=", 10.0/3)
	fmt.Println("10/3.0=", 10/3.0)
	fmt.Println("10%2=", 10%2)
	fmt.Println("10%3=", 10%3)
}
