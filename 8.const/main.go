package main

import "fmt"

//定数は型のない状態で宣言できる
//型のない状態だとintにもint64にも代入できる
const Pi = 1

var i1 int = Pi
var i2 int64 = Pi

const (
	URL      string = "https://xxx.co.jp"
	SiteName string = "Example"
)

//定数はすべて同じ値になる（使うことはなさそう）
const (
	A = "A"
	B
	C
)

//0から順番にインクリメントされる
const (
	a int = iota
	b
	c
)

func main() {
	fmt.Println(Pi)
	fmt.Println(URL)
	fmt.Println(SiteName)
	fmt.Println(A, B, C)
	fmt.Println(a, b, c)
}
