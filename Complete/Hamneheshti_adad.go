package main

import "fmt"

var (
	a, b int
	nums []int
)

func divi(a, b int) []int {
	var result []int
	for x := 2; x <= a-b; x++ {
		if (a-b)%x == 0 {
			result = append(result, x)
		}
	}

	return result
}

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		for _, j := range divi(a, b) {
			fmt.Print(j, " ")
		}
	} else {
		for _, j := range divi(b, a) {
			fmt.Print(j, " ")
		}
	}

}
