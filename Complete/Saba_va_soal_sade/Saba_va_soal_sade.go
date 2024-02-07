package main

import (
	"fmt"
	"math"
)

/*
	problem link : https://quera.org/problemset/31025
*/

func khosh(n, k float64) float64 {
	var i, res float64
	for i = 0; i < k; i++ {
		n = n / 2
	}
	res = n
	return res
}

func main() {
	var n, k float64
	fmt.Scan(&n, &k)

	fmt.Println(math.Floor(khosh(n, k)))
}
