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

// 関数の前にgoと付けるだけで平行処理できる
// 平行処理は並列処理と違い、一つのCPU上で順番関係なく処理する
// 並列処理よりメモリを節約できる
func main() {
	go sub()
	go sub()

	for i := 0; i < 10; i++ {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
