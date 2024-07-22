package main

import "fmt"

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 100)
	}
}

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don'n know")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don'n know")
	}

	switch n := 6; {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}

	var x interface{} = 3
	//型アサーションを使うことで型復元できる
	//この場合最初にint型で定義しているので、floatにはできずエラーになる
	i, isInt := x.(int)
	fmt.Println(i+2, isInt)
	//ただ変換時の返り値に2つ目を受け取った場合はエラーにならない
	f, isFloat := x.(float64)
	fmt.Println(f, isFloat)

	if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	} else {
		fmt.Println("I don't know")
	}

	//typeをswitchで判定できる
	switch v := x.(type) {
	case int:
		fmt.Println(v, "Int")
	case string:
		fmt.Println(v, "String")
	case bool:
		fmt.Println(v, "Boolean")
	default:
		fmt.Println(v, "I don't know")
	}

	//型アサーションによって一つの引数にいろんな型を入れれる
	anything("aaa")
	anything(123)
}
