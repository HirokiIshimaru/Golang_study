// defer
// 遅延実行
package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo") //hello fooの後に実行

	fmt.Println("hello foo")
}

func main() {
	// foo()

	// defer fmt.Println("world") //hello の後に実行

	// fmt.Println("hello")

	//スタッキングdefer
	// deferが続いた場合、後ろから前に実行される
	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/

	file, _ := os.Open("./practice05.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
