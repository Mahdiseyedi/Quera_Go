package main

import (
	"fmt"
	"math"
)

/*
	problem link : https://quera.org/problemset/3403
*/

func sum(a, b, c, d float64) float64 {
	res := a + b + c + d
	return res
}

func product(a, b, c, d float64) float64 {
	res := a * b * c * d
	return res
}

func max(a, b, c, d float64) float64 {
	res1 := math.Max(a, b)
	res2 := math.Max(res1, c)
	res3 := math.Max(res2, d)

	return res3
}

func min(a, b, c, d float64) float64 {
	res1 := math.Min(a, b)
	res2 := math.Min(res1, c)
	res3 := math.Min(res2, d)

	return res3
}

func avg(a, b, c, d float64) float64 {
	res := sum(a, b, c, d) / 4
	return res
}

func main() {
	var num1, num2, num3, num4 float64
	fmt.Scan(&num1, &num2, &num3, &num4)

	fmt.Printf("Sum : %.6f \n", sum(num1, num2, num3, num4))
	fmt.Printf("Average : %.6f \n", avg(num1, num2, num3, num4))
	fmt.Printf("Product : %.6f \n", product(num1, num2, num3, num4))
	fmt.Printf("MAX : %.6f \n", max(num1, num2, num3, num4))
	fmt.Printf("MIN : %.6f \n", min(num1, num2, num3, num4))
}
