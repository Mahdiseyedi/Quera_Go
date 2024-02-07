package main

import "fmt"

/*
	problem link : https://quera.org/problemset/34081
*/

var (
	a, b, c, d, e, f, result int
)

func slicer(nums []int, n int) []int {
	var aval, dovom, final []int
	aval = nums[len(nums)-n:]
	dovom = nums[0 : len(nums)-n]
	final = append(aval, dovom...)
	return final
}

func cardPakhshkon(n int) []int {
	var numbers []int
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	var stepTest []int
	stepTest = cardPakhshkon(a)
	step := 0
	for i := 0; i < 10000000000; i++ {
		stepTest = slicer(stepTest, b)
		if stepTest[0] == 1 {
			step = i + 1
			break
		}
	}

	fmt.Println(step)
}
