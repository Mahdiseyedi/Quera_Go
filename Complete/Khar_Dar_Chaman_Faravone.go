package main

import (
	"fmt"
)

func khar(a, b, l int) int {
	result := 0
	for i := 1; i < l+1; i++ {
		if i%2 == 0 {
			result += b
		} else {
			result += a
		}
	}

	return result
}

func main() {
	var a, b, l int
	fmt.Scan(&a, &b, &l)
	fmt.Println(khar(a, b, l))
}
