package main

import (
	"fmt"
	"strings"
)

/*
	problem link : https://quera.org/problemset/10327
*/

var (
	a, b, c, d, correct  string
	m, n, k, t           int
	nums                 []int
	names, corrent, test []string
	holy                 [][]string
	res                  bool
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Scan(&m)
	fmt.Scan(&correct)
	corrent = append(corrent, strings.Split(correct, "")...)
	for i := 0; i < m; i++ {
		fmt.Scan(&d)
		names = append(names, strings.Split(d, "")...)
		holy = append(holy, names)
		names = []string{}
	}

	for p := 0; p < m; p++ {
		res = true
		for q := 0; q < len(holy[p]); q++ {
			if !contains(corrent, holy[p][q]) {
				res = false
				continue
			}
		}
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
