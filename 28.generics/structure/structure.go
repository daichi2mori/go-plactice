package structure

import "fmt"

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

// メソッド定義時に型パラメータ指定は不可
//func (t T) m[U any][x U] {}

func Structure() {
	// Aはint
	// BはCのスライスかつAのポインタ
	// CはAのポインタ
	var t T[int, []*int, *int]
	fmt.Printf("A: %T, B: %T, C: %T", t.a, t.b, t.c)
}
