package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
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

	//2つのゴルーチンで作業を分散させながら、最後に合計をとっている
	s := []int{7, 2, 8, -9, 4, 0}
	go sum(s[:len(s)/2], ch1)
	go sum(s[len(s)/2:], ch1)
	x, y := <-ch1, <-ch1
	fmt.Println(x, y, x+y)

	//チャネル宣言時に容量を設定できる
	//容量を超える数を格納しようとするとエラーになる
	//取り出すときは、格納した順番で出てくる
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	//送受信は相手の準備が整うまでブロックされる
	//これによりゴルーチンが明示的なロックや条件変数なしで同期できます
	ch4 := make(chan int)
	ch5 := make(chan int)

	go reciever(ch4)
	go reciever(ch5)

	i := 0
	for i < 5 {
		ch4 <- i
		ch5 <- i
		time.Sleep(500 * time.Millisecond)
		i++
	}

	//チャネルでfor文を使うときはclose()しないとdeadlockエラーになる
	ch6 := make(chan int, 10)
	go fibonacci(cap(ch6), ch6)
	for i := range ch6 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}
