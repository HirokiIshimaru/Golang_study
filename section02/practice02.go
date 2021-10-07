// for文
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
			// continueが実行されると実行された時のforループが飛ばされる
		}
		if i > 5 {
			fmt.Println("break")
			break
			//	breakと通るとforループを抜ける
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// for {	//無限ループ
	// 	fmt.Println("hello")
	// }
}
