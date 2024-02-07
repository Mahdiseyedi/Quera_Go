package main

import "fmt"

/*
	problem link : https://quera.org/problemset/305
*/

var (
	a, b int
)

func hcf(n1 int, n2 int) int {
	if n2 != 0 {
		return hcf(n2, n1%n2)
	} else {
		return n1
	}
}

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(hcf(a, b))
}
