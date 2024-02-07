package main

import (
	"fmt"
	"strconv"
)

/*
	problem link : https://quera.org/problemset/595
                   https://quera.org/problemset/3410
*/

func fact(n, k int) int {
	n1 := 1
	n2 := 1
	n3 := 1

	for i := 0; i < n; i++ {
		n1 *= i + 1
	}
	for j := 0; j < k; j++ {
		n2 *= j + 1
	}
	for m := 0; m < n-k; m++ {
		n3 *= m + 1
	}
	res := n1 / (n2 * n3)
	return res
}

func line(n int) string {
	linee := ""
	for i := 0; i <= n-1; i++ {
		linee += strconv.Itoa(fact(n-1, i)) + " "
	}

	return linee
}

func khayam(n int) {
	for j := 0; j <= n; j++ {
		fmt.Print(line(j), "\n")
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	//fmt.Println(fact(5, 5))
	//fmt.Println(line(5))
	khayam(n)
}
