package main

import "fmt"

var (
	n, res, h, v int
)

func main() {

	fmt.Scan(&n)
	v = n / 2
	h = n - v
	res = (h + 1) * (v + 1)
	fmt.Println(res)
}
