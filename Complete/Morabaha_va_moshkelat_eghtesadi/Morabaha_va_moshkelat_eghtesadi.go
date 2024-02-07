package main

import (
	"fmt"
	"math"
)

/*
	problem link : https://quera.org/problemset/20249
*/

var (
	i, n, m, k, t, res, sum float64
	number                  []float64
)

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	sum = 0
	for i = 0; i < n; i++ {
		fmt.Scan(&k)
		sum += k
		number = append(number, k)
	}

	t = ((m * n) - sum) / m
	res = math.Floor(t)

	fmt.Println(res)
}
