package main

import (
	"fmt"
	"strings"
)

/*
	problem link : https://quera.org/problemset/6082
*/

func starCounter(inputs []string) int {
	var fati []string
	var res int

	// * = 42
	for j := 0; j < len(inputs); j++ {
		fati = strings.Split(inputs[j], "")
		for i := 0; i < len(fati); i++ {
			if fati[i] == "*" {
				res++
			}
		}
	}
	return res
}

func main() {

	var m, n int
	fmt.Scan(&m, &n)

	var line string
	var inputs []string

	for i := 0; i < m; i++ {
		fmt.Scan(&line)
		inputs = append(inputs, line)
	}

	fmt.Println(starCounter(inputs))
}
