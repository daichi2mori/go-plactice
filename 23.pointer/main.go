package main

import "fmt"

func main() {
	var n int = 100
	//変数の前に&をつけることで参照アドレスを表示する
	fmt.Println(&n)

	//基本型は値渡しをするのでpとnの値は共通されない
	var p int = n
	p = p * 2
	fmt.Println(n)

	//型の前にアスタリスクをつけることでポインタ型になる
	var p2 *int = &n
	fmt.Println(p2)

	//参照元を代入しているのでnとp2の値が共通
	*p2 = 200
	fmt.Println(n)
	n = 300
	fmt.Println(*p2)
}
