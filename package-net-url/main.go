package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.com/search?a=1&b=2#tope")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Scheme", u.Scheme)
	fmt.Println("Host", u.Host)
	fmt.Println("Path", u.Path)
	fmt.Println("RawQuery", u.RawQuery)
	fmt.Println("Fragment", u.Fragment)
	fmt.Println("Query", u.Query())

	fmt.Println("-----------------------------------------------")

	url := &url.URL{}
	url.Scheme = "https"
	url.Host = "example.com"
	q := url.Query()
	q.Set("q", "Go言語")

	url.RawQuery = q.Encode()

	fmt.Println(url)
}
