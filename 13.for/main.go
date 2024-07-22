package main

import "fmt"

func main() {
	i := 0
	for {
		if i == 3 {
			break
		}
		i++
		fmt.Println("for")
	}

	point := 0
	for point < 10 {
		fmt.Print(point)
		point++
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Print(i)
	}
	fmt.Println()

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()

	arr2 := [4]int{1, 2, 3, 4}
	for index, value := range arr2 {
		//valueが配列から取り出した値
		fmt.Println(index, value)
	}

	sl := []string{"Golang", "Rust", "Java"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		//Mapの場合はインデックスがキーになる
		fmt.Println(k, v)
	}

	//ラベル付きfor文
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理しない")
		}
		fmt.Println("処理しない")
	}
	fmt.Println("END")
}
