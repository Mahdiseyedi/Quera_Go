package main

import (
	"fmt"
)

func main() {

	var count int
	fmt.Scan(&count)

	var top, side string
	side += "*"
	for i := 0; i < count; i++ {
		top += "*"

	}
	for j := 1; j < count-1; j++ {
		side += " "
	}
	side += "*"

	var result string
	result += top
	result += "\n"
	for k := 0; k < count-2; k++ {
		result += side
		result += "\n"
	}
	result += top
	fmt.Println(result)
}
