package main

import (
	"fmt"
	"strings"
)

var (
	n, m, k, r, u              uint64
	x, y, z, f                 float64
	s, t, l, g, v, result, res string
	nums1, nums2               []int
	names1, names2             []string
)

func summer(pp, qq string) string {
	res = ""
	v = ""
	if len(pp) == len(qq) {
		res += "2"
		for i := 1; i < len(pp); i++ {
			v += "0"
		}
		res += v
	} else if len(pp) > len(qq) {
		res += pp[0 : len(pp)-len(qq)]
		res += qq
	} else if len(qq) > len(pp) {
		res += qq[0 : len(qq)-len(pp)]
		res += pp
	}

	return res
}

func multiple(pp, qq string) string {
	result = "1"
	for i := 0; i < len(pp)+len(qq)-2; i++ {
		result += "0"
	}

	return result
}

func main() {
	fmt.Scan(&s)
	fmt.Scan(&t)
	fmt.Scan(&l)

	if strings.TrimSpace(t) == "+" {
		fmt.Println(summer(s, l))
	} else {
		fmt.Println(multiple(s, l))
	}
}
