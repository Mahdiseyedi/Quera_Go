package main

import "fmt"

/*
	problem link : https://quera.org/problemset/3029
*/

var (
	n, m, k, t, res int
)

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	fmt.Scan(&t)

	if k >= n {
		fmt.Println("Right")
	} else {
		fmt.Println("Left")
	}
}
