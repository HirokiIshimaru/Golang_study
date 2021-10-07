// if文
package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	// num := 6
	// if num%2 == 0 { //numを2で割りあまりが0だった場合の処理
	// 	fmt.Println("by 2")
	// } else if num%3 == 0 {
	// 	fmt.Println("by 3")
	// } else { //上記の条件に引っ掛からなかった場合の処理
	// 	fmt.Println("else")
	// }

	// &&
	x, y := 1, 10
	if x == 10 && y == 10 { //xとyが両方とも10だった場合の処理
		fmt.Println("&&")
	}

	// ||
	if x == 10 || y == 10 { //xとyの片方が10だった場合の処理
		fmt.Println("||")
	}
}
