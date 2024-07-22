package main

import "fmt"

func main() {
	a := 10
	b := 20

	// 算術演算子
	fmt.Println("a + b =", a+b) // 30
	fmt.Println("a - b =", a-b) // -10
	fmt.Println("a * b =", a*b) // 200
	fmt.Println("a / b =", a/b) // 0
	fmt.Println("a % b =", a%b) // 10

	// 代入演算子
	a += b
	fmt.Println("a += b -> a =", a) // 30
	a -= b
	fmt.Println("a -= b -> a =", a) // 10

	// 比較演算子
	fmt.Println("a == b =", a == b) // false
	fmt.Println("a != b =", a != b) // true
	fmt.Println("a < b =", a < b)   // true
	fmt.Println("a <= b =", a < b)  // true
	fmt.Println("a > b =", a < b)   // false
	fmt.Println("a >= b =", a < b)  // false

	// 論理演算子
	fmt.Println("true && false =", true && false) // false
	fmt.Println("true || false =", true || false) // true
	fmt.Println("!true =", !true)                 // false

	// ビット演算子
	fmt.Println("a & b =", a&b)   // 0
	fmt.Println("a | b =", a|b)   // 30
	fmt.Println("a ^ b =", a^b)   // 30
	fmt.Println("a << 1 =", a<<1) // 20
	fmt.Println("a >> 1 =", a>>1) // 5

	// インクリメントとデクリメント
	a++
	fmt.Println("a++ -> a =", a) // 11
	a--
	fmt.Println("a-- -> a =", a) // 10

	// ポインタ演算子
	p := &a
	fmt.Println("&a =", p)  // アドレス
	fmt.Println("*p =", *p) // 10
}
