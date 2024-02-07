package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/91712
*/

var (
	x1, x2, h, m int
	res          = ""
)

func main() {

	fmt.Scan(&x1)
	fmt.Scan(&x2)

	if x1 == x2 {
		res = "Saal Noo Mobarak!"
	} else if x1 > x2 {
		for i := 0; i < x1-x2; i++ {
			res += "L"
		}
	} else {
		for i := 0; i < x2-x1; i++ {
			res += "R"
		}
	}

	fmt.Println(res)
}
