package main

import (
	"fmt"
	"strconv"
)

/*
	problem link : https://quera.org/problemset/10325
*/

func main() {

	var r, c int
	fmt.Scan(&r, &c)
	var result string

	if c <= 10 {
		result += "Right "
		r = 11 - r
		result += strconv.Itoa(r)
		result += " "
		result += strconv.Itoa(c)
	} else {
		result += "Left "
		c = c - 10 + 1
		r = 11 - r
		result += strconv.Itoa(r)
		result += " "
		result += strconv.Itoa(c)
	}

	fmt.Println(result)
}
