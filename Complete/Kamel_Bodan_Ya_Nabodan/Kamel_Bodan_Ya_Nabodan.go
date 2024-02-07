package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/282
*/

func main() {

	var a int
	fmt.Scan(&a)

	result := 0

	for i := 1; i < a; i++ {
		if a%i == 0 {
			result += i
		}
	}
	if result == a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
