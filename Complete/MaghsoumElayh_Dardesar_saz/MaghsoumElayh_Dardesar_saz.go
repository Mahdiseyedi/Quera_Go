package main

import "fmt"

/*
	problem link : https://quera.org/problemset/33045
*/

func vasat(m int) []int {
	var res []int

	for i := 1; i <= m; i++ {
		if m%i == 0 {
			res = append(res, i)
		}
	}

	return res
}

func counter(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res++
	}

	return res
}

func sumer(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	count := 0

	for i := 0; i <= n; i++ {
		sum += sumer(vasat(i))
		count += counter(vasat(i))
	}

	fmt.Print(count, " ", sum)
}
