// range
package main

import (
	"fmt"
)

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	// ↑の処理を簡潔に書く(range)

	for i, v := range l {
		fmt.Println(i, v)
	}
	// iを使いたくない場合
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m { //keyをvalueを表示
		fmt.Println(k, v)
	}
	for k := range m { //keyのみ表示
		fmt.Println(k)
	}
	for _, v := range m { //valueのみ表示
		fmt.Println(v)
	}
}
