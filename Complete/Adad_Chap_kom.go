package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compare(num string) {

	var c int
	var ts string
	result := strings.Split(num, "")
	for i := 0; i < len(result); i++ {
		c, _ = strconv.Atoi(result[i])
		ts = strconv.Itoa(c) + ": "
		for j := 0; j < c; j++ {
			ts += strconv.Itoa(c)
		}
		fmt.Println(ts)
	}
}

func main() {

	var a string
	fmt.Scan(&a)
	compare(a)
}
