package main

import "fmt"

//このエラーハンドリングは使わないほうが良い
func main() {
	//recoverはpanic状態のとき返り値が発生する
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	//panicだけだと実行中のプログラムを終了させる
	panic("runtime error")
	fmt.Println("start")

	//recoverはpanicによって発生したエラーから復帰させる
	//recoverは基本的にdefer内で使用する
}
