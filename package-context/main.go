package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ch chan string) {
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
	ch <- "execution result"
}

func main() {
	ch := make(chan string)

	ctx := context.Background()

	// WithTimeoutで設定した時間で処理が終了しなかったら、
	// select文の case <-ctx.Done() に入りエラーになる
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	go longProcess(ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("#####Error#####")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}

	fmt.Println("out of loop")
}
