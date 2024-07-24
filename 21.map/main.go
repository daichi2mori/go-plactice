package main

import "fmt"

func main() {
	//map[キーの型]を設定
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"C": 300, "D": 400}
	fmt.Println(m2)

	//make関数は動的なメモリを確保するために使用されるが、
	//初期化のための値を渡すことはできない
	m3 := make(map[int]string)
	m3[1] = "A"
	m3[2] = "B"
	fmt.Println(m3)

	//map関数から値を取り出したときに成功したかの判定ができる
	s, ok := m3[4]
	if !ok {
		fmt.Println("値無し")
	}
	fmt.Println(s, ok)

	//delete関数でmapから要素を削除できる
	//第一引数にはmap、第二引数には削除したいキー
	delete(m3, 1)
	fmt.Println(m3)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
