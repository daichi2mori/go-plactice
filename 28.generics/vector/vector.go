package vector

import "fmt"

type vector[T any] []T

type intVector = vector[int]

func Vector() {
	var v vector[int] = vector[int]{10, 20, 30}
	fmt.Println(v)

	var v2 vector[float64] = vector[float64]{1.3, 3.4, 5.6}
	fmt.Println(v2)

	// type作成時点でインスタンス化しているためそのまま使える
	v3 := intVector{1, 2, 3}
	fmt.Println(v3)
}
