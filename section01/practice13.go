package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}

func add(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}
func Hello(name string) {
	fmt.Println(name)
}

func main() {
	// 関数
	// 関数helloを呼び出し
	hello()
	Hello("Tanaka")

	// add(10, 20) //関数addを呼び出し
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
}
