package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	//送信受信のどちらもできる
	//以下の記述だとnilになるのでmake関数を使う必要がある
	//var ch1 chan int
	//ch1 = make(chan int)

	//これだと1行ですむ
	ch1 := make(chan int)

	//受信専用
	//var ch2 <-chan int

	//受信専用
	//var ch3 chan<- int

	s := []int{7, 2, 8, -9, 4, 0}
	//送受信は相手の準備が整うまでブロックされる
	//これにより、ゴルーチンが明示的なロックや条件変数なしで同期できます
	go sum(s[:len(s)/2], ch1)
	go sum(s[len(s)/2:], ch1)
	x, y := <-ch1, <-ch1

	//2つのゴルーチンで作業を分散させながら、最後に合計をとっている
	fmt.Println(x, y, x+y)
}
