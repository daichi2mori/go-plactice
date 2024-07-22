package main

import "fmt"

func main() {
	var fl64 float64 = 6.4
	fmt.Println((fl64))

	//暗黙的な型定義だと自動でfloat64になる
	//これは環境依存しない
	fl := 5.4
	fmt.Printf("%T\n", fl)
	fmt.Println(fl64 + fl)

	var fl32 float32 = 3.2
	fmt.Printf("%T\n", fl32)
	fmt.Printf("%T\n", float64(fl32))

	//floatは無限大やnot a numberを保持することができる
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	var u8 uint = 255
	fmt.Printf("%T\n", u8)
	/*
		uintは正の整数という型
		uintにもuint8,uint16,uint32,uint64の型がある
	*/

	var c64 complex64 = -5 + 12i
	fmt.Println(c64)
	fmt.Printf("%T\n", c64)
	/*
		complexは複素数型
		complex32,complex64の2種類がある
	*/
}
