package main

import (
	"fmt"
)

func fibo() []uint64 {
	var n1, n2, i uint64
	n1 = 1
	n2 = 2
	var result []uint64
	result = append(result, n1, n2)

	for i = 1; i <= 100; i++ {
		result = append(result, result[i]+result[i-1])
	}
	return result
}

func contains(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func target(n uint64) string {
	var i uint64
	final := "+"
	for i = 2; i <= n; i++ {
		if contains(fibo(), i) {
			final += "+"
		} else {
			final += "-"
		}
	}
	return final
}

func main() {
	var n uint64
	fmt.Scan(&n)
	fmt.Println(target(n))
}
