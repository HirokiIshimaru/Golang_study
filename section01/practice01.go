package main

// Golangのimportのpackageはここに書く
import (
	"fmt"
	"os/user"
	"time"
)

// initは一番初めに表示される関数(main関数より先に表示される)
// 初期設定によく使われる
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
}
