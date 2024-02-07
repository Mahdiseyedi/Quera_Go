package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/17244
*/

func main() {
	var num int
	fmt.Scan(&num)

	res := (num * (num + 1)) / 2
	fmt.Println(res)
}
