package main

import "fmt"

func foo(num ...int) {
	fmt.Println(len(num), num)
}
func main() {
	// 可変長引数
	foo(10, 20, 30)
}
