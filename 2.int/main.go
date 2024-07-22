package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i + 50)

	/*
		intには int8,int16,int32,int64がある。
		型指定でintだけだと環境に依存したビット数になる。
		（Windowsの64bit版だとint64になる）
	*/

	var i2 int64 = 200
	/*
		明示的にintのビット数を指定した変数と、
		ただのint型だけ指定した変数での計算はできない
	*/
	//fmt.Println(i + i2)

	//"%T"で型情報を表示できる
	fmt.Printf("%T\n", i2)

	//型変換 int64で定義した変数をint32に変換
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}
