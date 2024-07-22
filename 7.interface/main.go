package main

import "fmt"

func main() {
	/*
		interfaceは初期値にnilをもっている
		nilは値が存在しないことを表している
		nullとは違いnilは型を持っている
		interfaceはすべての型と互換性を持っている
	*/
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = "string"
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	//計算はできない
	// x = 2
	// fmt.Println(x + 3)
}
