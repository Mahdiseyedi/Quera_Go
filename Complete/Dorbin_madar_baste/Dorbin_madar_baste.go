package main

import "fmt"

/*
	problem link : https://quera.org/problemset/2794
*/

func junk(n, m, k int) int {
	l := 0
	if n == m {
		l = k
	} else if n == k {
		l = m
	} else if m == k {
		l = n
	}

	return l
}

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	fmt.Print(junk(a, c, e), " ", junk(b, d, f))

}
