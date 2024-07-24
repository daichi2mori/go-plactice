package main

import "fmt"

func main() {
	//配列は[]に配列数を入れるが、スライスは何も入れない
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"a", "b"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[1])
	fmt.Println(sl5[2:4]) //配列数4番目の手前まで
	fmt.Println(sl5[:2])  //最初から2番目まで
	fmt.Println(sl5[2:])  //2番目から最後まで

	//sliceはappendを使うことで要素数を増やせる
	sl2 = append(sl2, 300)
	fmt.Println(sl2)

	//複数設定できる
	sl2 = append(sl2, 400, 500, 600)
	fmt.Println(sl2)

	//make関数の第二引数は要素数、第三引数は容量を設定できる
	sl6 := make([]int, 5, 8)
	fmt.Println(len(sl6), cap(sl6))
	//容量を要素数より大きく設定しておくと、
	//メモリを連続的かつ先に確保し、メモリ管理が効率的になる

	//可変長for
	fmt.Println(Sum(1, 2, 3))
	//配列を渡すときは最後にスプレッド演算子を付ける
	fmt.Println(Sum(sl2...))

	//固定長配列でスプレッド演算子
	//arr[:]で固定長配列をスライスに変換できる
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Sum(arr[:]...))
}

//引数をスプレッド演算子にすることで可変な配列を処理できる
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
