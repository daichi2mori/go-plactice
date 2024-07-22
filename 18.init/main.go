package main

import "fmt"

//init関数がある場合はmain関数より先に実行される
//複数定義することもでき、その場合は上から順番に実行される（意味はないが）
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}

func main() {
	fmt.Println("Main")
}
