package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func main() {
	i := []int{4, 9, 1, 3, 2, 5, 8, 7, 6}
	s := []string{"j", "1", "a"}

	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)

	fmt.Println("------------------------------------------------------------")

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	sort.Slice(el, func(i, j int) bool {
		fmt.Println("i =", i, "  j =", j)
		return el[i].Name < el[j].Name
	})
	fmt.Println("------------------------------------------------------------")
	fmt.Println(el)
}
