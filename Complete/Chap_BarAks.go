package main

import (
	"fmt"
)

func main() {
	// var nums []string
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')

	// result := strings.TrimSpace(input)
	// nums = strings.Split(result, "\n")

	// for i := len(nums)  i > 0  i-- {
	// 	fmt.Println(nums[i])
	// }
	var num int
	fmt.Scan(&num)
	var arr [1000]int
	var i int = 0
	for num != 0 {
		arr[i] = num
		fmt.Scan(&num)
		i++
	}
	for j := i - 1; j >= 0; j-- {
		fmt.Print(arr[j])
		fmt.Printf("\n")
	}
}
