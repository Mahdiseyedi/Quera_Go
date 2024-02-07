package main

import (
	"fmt"
	"strconv"
)

/*
	problem link : https://quera.org/problemset/649
*/

func isPrime(n int) bool {
	res := true
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			res = false
			break
		}
	}
	if n == 1 {
		res = false
	}

	return res
}

func primePrinter(a, b int) string {
	result := ""
	for i := a; i < b; i++ {
		if isPrime(i) && i != a && i != b {
			result += strconv.Itoa(i) + ","
		}
	}

	if len(result) > 1 {
		result = result[0 : len(result)-1]
	}

	return result
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(primePrinter(a, b))
}
