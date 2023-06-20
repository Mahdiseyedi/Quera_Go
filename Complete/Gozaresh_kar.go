package main

import "fmt"

func check(values []int, k int) bool {
	res := false
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	if total >= k {
		res = true
	}

	return res
}

func main() {
	var n, k, t int
	fmt.Scan(&n, &k)

	var values []int

	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		values = append(values, t)
	}

	if check(values, k) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
