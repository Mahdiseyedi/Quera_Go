package main

import "fmt"

var (
	n, m, k, turn, res int
	result, name, step string
	numbers            []int
	names              []string
)

func main() {
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		numbers = append(numbers, m)
	}

	turn = 0
	for j := 0; j < n-1; j++ {
		if numbers[j] != numbers[j+1] {
			turn++
		}
	}
	fmt.Println(turn)
}
