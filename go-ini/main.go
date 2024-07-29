package main

import (
	"fmt"
	"go-ini/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.SQLDriver)
}
