package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/51865
*/

func mark(x, n int) int {
	result := 0
	if n == 0 {
		result = 20
	} else if n == 7 {
		result = x
	} else {
		result = x - n
	}

	if result < 0 {
		result = 0
	}

	return result
}

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	fmt.Println(mark(x, n))
}
