package main

import (
	"fmt"
	"strconv"
)

var (
	i, n, m, k, t, sum int
	res, final         string
	number             []int
)

func khar(m, n int) string {
	res = ""
	if m == n || m == n+2 {
		if m%2 == 0 && n%2 == 0 {
			res += strconv.Itoa(m + n)
		} else {
			res += strconv.Itoa(m + n - 1)
		}

	} else {
		res += "-1"
	}
	return res
}

func main() {
	fmt.Scan(&t)
	for i := 0; i < 2*t; i++ {
		fmt.Scan(&sum)
		number = append(number, sum)
	}
	for j := 0; j < len(number); j += 2 {
		fmt.Println(khar(number[j], number[j+1]))
	}
}
