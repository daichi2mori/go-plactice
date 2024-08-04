package set

import "fmt"

// comparableは比較可能な型（int,stringなどの基本型）
type S[T comparable] map[T]struct{}

func newSet[T comparable](xs ...T) S[T] {
	s := make(S[T])
	for _, v := range xs {
		s.Add(v)
	}
	return s
}

func (s S[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s S[T]) Remove(x T) {
	delete(s, x)
}

func Set() {
	s := newSet(1, 2, 3)
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)
}
