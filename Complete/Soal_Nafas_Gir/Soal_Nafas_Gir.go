package main

import "fmt"

/*
	problem link : https://quera.org/problemset/26651
*/

var (
	n, m, k, r, u  int
	s, t, l        string
	nums1, nums2   []int
	names1, names2 []string
)

func main() {
	fmt.Scan(&m)

	for i := 0; i < 2*m; i++ {
		if i < m {
			fmt.Scan(&n)
			nums1 = append(nums1, n)
		} else {
			fmt.Scan(&n)
			nums2 = append(nums2, n)
		}
	}

	res := 0
	for l := 0; l < m; l++ {
		res += nums1[l] * nums2[l]
	}

	fmt.Println(res)
}
