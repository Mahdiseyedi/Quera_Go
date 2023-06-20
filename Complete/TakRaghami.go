package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compare(num string) string {

	for len(num) > 1 {
		var b, c int
		result := strings.Split(num, "")
		for i := 0; i < len(result); i++ {
			c, _ = strconv.Atoi(result[i])
			b += c
		}
		num = strconv.Itoa(b)
	}
	return num
}

func main() {

	var a string
	fmt.Scan(&a)
	fmt.Println(compare(a))
}
