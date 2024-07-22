package main

import "fmt"

func main() {
	/*
		Go言語では最初に宣言した配列数を変更することはできない
		Printfで型表示すると[3]intとなっており、
		要素数3までが型として定義されているため
	*/
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//[...]は宣言時の要素数を自動で算出してくれる
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[1])

	arr2[2] = "C"
	fmt.Println(arr2)

	//要素数を調べる
	fmt.Println(len(arr1))
}
