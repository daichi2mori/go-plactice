package typesets

import "fmt"

// ~(チルダ)をつけることで独自設定した型の、元の型を判別してくれる
type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func max[T Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type MyInt int

func PrintNumber() {
	fmt.Println(max[int8](1, 2))
	fmt.Println(max[int32](5, 2))
	fmt.Println(max[float32](1.6, 2.5))
	fmt.Println(max[MyInt](8, 9))
}
