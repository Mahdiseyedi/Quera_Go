package main

import (
	"fmt"
	"strings"
)

func contains(s []string, e string) int {
	count := 0
	for _, a := range s {
		if a == e {
			count++
		}
	}
	return count
}

func maghloub(input []string, number string) string {
	result := ""
	test := ""
	for i := len(input); i > 0; i-- {
		result += input[i-1]
	}
	if result == number {
		test = "YES"
	} else {
		test = "NO"
	}

	return test
}

func main() {
	var s string
	fmt.Scan(&s)
	test := strings.Split(s, "")
	fmt.Println(maghloub(test, s))
}
