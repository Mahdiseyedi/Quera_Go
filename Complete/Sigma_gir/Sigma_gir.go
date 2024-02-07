package main

import "fmt"

/*
	problem link : https://quera.org/problemset/647
*/

var (
	n, m, k, r, u  int
	x, y, z, f     float64
	s, t, l        string
	nums1, nums2   []int
	names1, names2 []string
	res, result    bool
)

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	k = 0
	r = 0
	for i := -10; i <= m; i++ {
		for j := 1; j <= n; j++ {
			r += (i + j) * (i + j) * (i + j) / (j * j)
		}
		k += r
		r = 0
	}

	fmt.Println(k)
}
