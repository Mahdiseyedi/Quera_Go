package main

import (
	"fmt"
	"strings"
)

var (
	fnums                 []float64
	a, b, c, v, k, res    int
	ramz, g, h            string
	firstLine, secondLine []string
)

func contain(name, n string) bool {
	test := strings.Split(name, "")
	for _, i := range test {
		if i == n {
			return true
		}
	}

	return false
}

func main() {
	fmt.Scan(&k)
	fmt.Scan(&ramz)

	res := 0
	for i := 0; i < k; i++ {
		fmt.Scan(&g)
		if contain(g, string(ramz[i])) && i <= k/2 {
			res += i
		} else if contain(g, string(ramz[i])) {
			res += i - k
		}
	}
	fmt.Println(res)
}
