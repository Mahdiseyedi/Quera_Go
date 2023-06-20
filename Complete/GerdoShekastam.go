package main

import "fmt"

func main() {

	var n, x, y int
	fmt.Scan(&n, &x, &y)

	flag := false

	for i := 1; i <= 100000; i++ {

		if (n-(i*x))%y == 0 {
			flag = true
			j := (n - (i * x)) / y
			if i >= 0 && j >= 0 {
				fmt.Print(i, " ", j, "\n")
				break
			} else {
				fmt.Println(-1)
				break
			}
		}
	}
	if !flag {
		fmt.Println(-1)
	}

}
