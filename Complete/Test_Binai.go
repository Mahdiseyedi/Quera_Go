package main

import (
	"fmt"
	"strings"
)

func Bina(first, second string) int {
	res := 0

	oneLine := strings.Split(first, "")
	twoLine := strings.Split(second, "")

	for i := 0; i < len(oneLine); i++ {
		if oneLine[i] != twoLine[i] {
			res++
		}
	}

	return res
}

func main() {

	var num, first, second string
	fmt.Scan(&num, &first, &second)
	fmt.Println(Bina(first, second))
}
