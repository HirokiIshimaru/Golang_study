// log

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("test")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test") //Fatalを使うとコードが終了され、それ以降のコードは読まれない
	log.Fatalln("error!!")
	fmt.Println("ok!")
}
