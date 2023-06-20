package main

import "fmt"

func fisaChecker(x, y, z int) bool {
	return x*x+y*y == z*z || x*x == y*y+z*z || y*y == x*x+z*z
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if fisaChecker(a, b, c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
