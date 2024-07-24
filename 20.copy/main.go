package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000
	//スライスは参照型のため、そのままコピーして値を変更すると元の変数の値も変わる
	fmt.Println(sl)

	//コピー元のスライス
	a := []int{1, 2, 3, 4, 5}
	//コピー先のスライス
	b := make([]int, len(a))
	//copy関数を使うことでコピー元の変更を防げる
	copy(b, a)
	b[0] = 6
	fmt.Println(b)
}
