package main

import "fmt"

/*
	problem link : https://quera.org/problemset/28947
*/

func main() {
	var s int
	fmt.Scan(&s)
	total := []string{"seyed", "samano", "sib", "seke", "sajad", "sara", "sarava", "sebtin", "solo"}

	for i := 0; i < s; i++ {
		fmt.Println(total[i])
	}
}
