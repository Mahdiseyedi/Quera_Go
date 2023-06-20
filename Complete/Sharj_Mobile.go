package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	res := (num * (num + 1)) / 2
	fmt.Println(res)
}
