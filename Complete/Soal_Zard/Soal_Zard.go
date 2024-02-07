package main

import "fmt"

/*
	problem link : https://quera.org/problemset/3537
*/

func wowCreator(n int) string {
	var result = "W"
	for i := 1; i <= n; i++ {
		result += "o"
	}
	result += "w!"
	return result
}

func main() {
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println(wowCreator(input))
}
