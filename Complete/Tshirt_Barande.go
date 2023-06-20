package main

import "fmt"

var (
	n, m, k, r, u  int
	s, t, l        string
	nums1, nums2   []int
	names1, names2 []string
	res, result    bool
)

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	fmt.Scan(&r)

	if n >= k && m >= r {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
