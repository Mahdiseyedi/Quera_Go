package main

import "fmt"

func main() {
	var input int
	var i int
	var result = 1
	fmt.Scanf("%d", &input)
	for i = 1; i <= input; i++ {
		result *= i
	}
	fmt.Println(result)
}
