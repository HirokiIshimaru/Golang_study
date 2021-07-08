package main

import "fmt"

func main() {
	// map
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)          //全て出力
	fmt.Println(m["apple"]) //appleだけ出力
	m["banana"] = 300       //bananaの数値を変更
	fmt.Println(m)
	m["new"] = 500 //新しいデータを入れる
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	m2 := make(map[string]int) //空のmapを作成
	m2["pc"] = 5000            //空のmapに値を代入する
	fmt.Println(m2)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
