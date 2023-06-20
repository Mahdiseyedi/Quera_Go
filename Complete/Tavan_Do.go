package main

import (
	"fmt"
)

func compare(num int) int {
	var result int = 1
	for num >= result {
		result *= 2
	}
	return result
}

func main() {

	var a int
	fmt.Scan(&a)
	fmt.Println(compare(a))
}
