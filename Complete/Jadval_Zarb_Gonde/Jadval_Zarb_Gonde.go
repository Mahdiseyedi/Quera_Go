package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/3409
*/

func compare(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println(" ")
	}
}

func main() {

	var a int
	fmt.Scan(&a)
	compare(a)
}
