package main

import (
	"fmt"
)

func main() {

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a == c {
		fmt.Println("Vertical")
	} else if b == d {
		fmt.Println("Horizontal")
	} else {
		fmt.Println("Try again")
	}
}
