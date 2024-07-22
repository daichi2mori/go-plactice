package main

import "fmt"

func main() {
	//byte型をそのまま配列として宣言することができる
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	//格納時に文字列をバイト変換できる
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
