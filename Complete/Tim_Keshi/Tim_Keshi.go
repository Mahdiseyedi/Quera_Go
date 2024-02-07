package main

import (
	"fmt"
	"math"
)

/*
	problem link : https://quera.org/problemset/80651
*/

var (
	a, b, c, d, e, f, result float64
)

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Scan(&e)
	fmt.Scan(&f)

	result += math.Min(a, b)
	result += math.Min(c, d)
	result += math.Min(e, f)

	fmt.Println(result)
}
