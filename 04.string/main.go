package main

import "fmt"

func main() {
	/*
		string型はバイト配列の集まりであり、一文字だけ取り出すことができる。
		ただし表示される値はバイト値になるため、文字列を表示させたい場合は型変換が必要
	*/
	var s string = "Hello Golang"
	fmt.Println(s[0])
	fmt.Println(string(s[0]))

	//バッククォートを使うと改行して表示することもできる
	fmt.Println(`test
	test
		test`)

	//ダブルクォーテーションを表示させたいとき
	fmt.Println("\"")
	fmt.Println(`"`)
}
