package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	//GET
	res, _ := http.Get("https://example.com")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])
	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(string(body))

	fmt.Println("---------------------------------------")

	//POST
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	// res, _ = http.PostForm("https://example.com", vs)
	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(string(body))
}
