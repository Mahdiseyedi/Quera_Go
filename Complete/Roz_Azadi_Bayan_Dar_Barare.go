package main

import (
	"fmt"
)

// func compare(num int) {
// 	for i := 1; i <= num; i++ {
// 		for j := 1; j <= num; j++ {
// 			fmt.Print(i*j, " ")
// 		}
// 		fmt.Println(" ")
// 	}
// }

func main() {

	var a int
	fmt.Scan(&a)

	if a%2 != 0 {
		fmt.Println("Payin Barare")
	} else {
		fmt.Println("Bala Barare")
	}

}
