package main

import (
	"fmt"
	"strconv"

	"28.generics/set"
	"28.generics/structure"
	"28.generics/typesets"
	"28.generics/vector"
)

func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func PrintSliceStrings(s []string) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// Genericsを使うことで、上記の型だけが違う同じ処理を、
// 型安全に一つの関数で定義できる
// anyはinterface{}と同じで、そのままinterface{}と記述しても大丈夫
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type MyInt int

func (i MyInt) String() string {
	// MyInt型をint型に変換してから、string型に変換
	return strconv.Itoa(int(i))
}

// String()メソッドを持ったインターフェースのみ引数に取る
func PrintStringer[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

func main() {
	// PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSliceStrings([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})
	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})
	fmt.Println(PrintStringer([]MyInt{1, 2, 3, 4}))
	fmt.Println("-----typesets-----")
	typesets.PrintNumber()
	fmt.Println("-----vector-----")
	vector.Vector()
	fmt.Println("-----structure-----")
	structure.Structure()
	fmt.Println("-----set-----")
	set.Set()
}
