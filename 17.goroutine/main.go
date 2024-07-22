package main

import (
	"fmt"
	"time"
)

func sub() {
	for i := 0; i < 20; i++ {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

// 関数の前にgoと付けるだけで並列処理できる
func main() {
	go sub()
	go sub()

	for i := 0; i < 10; i++ {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
