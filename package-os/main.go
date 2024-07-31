package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 読み込みonly
	f1, err1 := os.Open("text.txt")
	if err1 != nil {
		log.Fatalln(err1)
	}
	defer f1.Close()

	//読み込み
	data1 := make([]byte, 1024)
	//countには読み込んだバイト数
	count, err2 := f1.Read(data1)
	if err2 != nil {
		log.Fatalln(err2)
	}
	fmt.Println(count)
	fmt.Println(string(data1))

	//書き込み権限付き
	//Create()はすでに同じ名前のファイルがあれば削除され、新しく作成される
	f2, err3 := os.Create("text.txt")
	if err3 != nil {
		log.Fatalln(err3)
	}
	data2 := []byte("Hello\n")
	data3 := []byte("Golang")
	f2.Write(data2)
	f2.Write(data3)
}
