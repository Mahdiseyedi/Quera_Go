package main

import (
	"fmt"
	"strconv"
)

var (
	k, t, h, m int
	res        = ""
)

func main() {

	fmt.Scan(&k)
	fmt.Scan(&t)

	h = 12 - k
	m = 60 - t

	if h >= 10 && h != 12 {
		res += strconv.Itoa(h)
	} else if h < 10 {
		res += "0"
		res += strconv.Itoa(h)
	} else {
		res += "00"
	}

	res += ":"

	if m >= 10 && m != 60 {
		res += strconv.Itoa(m)
	} else if m < 10 {
		res += "0"
		res += strconv.Itoa(m)
	} else {
		res += "00"
	}

	fmt.Println(res)
}
