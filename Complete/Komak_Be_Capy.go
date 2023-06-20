package main

import (
	"fmt"
)

func capy(n int, s string) string {

	result := ""
	copyText := "copy of "
	for i := 0; i < n; i++ {
		result += copyText
	}
	result += s

	return result
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	fmt.Println(capy(n, s))
}
