package main

import (
	"fmt"
	"sync"
	"time"
)

/*
複数のゴルーチンが同時に実行されると、共有リソースでの競合が発生する
それを避けるために mutex と呼ばれる排他制御機構がある
*/

// ミューテックスを保持する変数
var mutex *sync.Mutex
var waitGroup *sync.WaitGroup

var st struct {
	A int
	B int
	C int
}

// syncを使用しないと、stのそれぞれの値がバラバラになる
func UpdateAndPrint(n int) {
	mutex.Lock()
	defer mutex.Unlock()

	st.A = n
	st.B = n
	st.C = n
	fmt.Println(st)
	time.Sleep(time.Microsecond)
}

func main() {
	//ポインタは初期化しないとnilになる
	mutex = new(sync.Mutex)
	waitGroup = new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		//Add()でゴルーチンを追加
		waitGroup.Add(1)
		go func() {
			//ゴルーチンが完了したことをWait()に伝える
			defer waitGroup.Done()
			for i := 0; i < 100; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	//ゴルーチンを終了させる
	waitGroup.Wait()
}
