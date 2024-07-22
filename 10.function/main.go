package main

import "fmt"

/*
func 関数名 (引数 型) 型 {
    // do something
    return 値
}

func 関数名 (引数 型) (型, 型) {
    // do something
    return 値, 値
}
*/

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//返り値の変数と型を同時に設定できる
//それによってreturnの変数を省略できる
func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
}

//関数を返す
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

//引数に関数を使う
func CallFunction(f func()) {
	f()
}

//クロージャー
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

//ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 3)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()

	/*
		無名関数
	*/
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 3)
	fmt.Println(i5)

	i6 := func(x, y int) int {
		return x + y
	}(1, 3)
	fmt.Println(i6)

	f2 := ReturnFunc()
	f2()

	CallFunction(f2)

	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("my"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is"))
	fmt.Println(f3("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}
