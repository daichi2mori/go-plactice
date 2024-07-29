package main

import (
	"fmt"
	"os"
)

// deferは関数の終わりに実行される（return直後）
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

// deferはstack形式で実行される(6から起動)
// 後入れ先出し（Last In First Out : LIFO）
func RunDefer() {
	defer fmt.Println("4")
	defer fmt.Println("5")
	defer fmt.Println("6")
}

func main() {
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()

	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}
