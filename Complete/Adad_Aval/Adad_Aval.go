package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/293
*/

func isPrime(number int) bool {
	result := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			result = false
		}
	}
	if number == 1 {
		result = false
	}
	return result
}

func primeCounter(num1, num2 int) {
	for i := num1; i <= num2; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	var num1, num2 int

	fmt.Scan(&num1, &num2)
	primeCounter(num1, num2)

}
