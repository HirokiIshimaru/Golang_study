// switch
package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "test"
}

func main() {
	// os := "mac"
	os := getOsName()
	switch os {
	case "mac": //macだった場合の処理
		fmt.Println("mac!!")
	case "windows": //windowsだった場合の処理
		fmt.Println("windows!!")
	default: //defaultだった場合の処理
		fmt.Println("default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Mourning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
