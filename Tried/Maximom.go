package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func capy(num string, nums2 []string) int {
	num += "mahdi"
	max, _ := strconv.Atoi(nums2[0])
	var turn int
	for i := 1; i < len(nums2); i++ {
		turn, _ = strconv.Atoi(nums2[i])
		if turn > max {
			max = turn
		}
	}

	return max
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	nums := strings.Split(input, " ")

	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	result := strings.TrimSpace(input2)
	nums2 := strings.Split(result, " ")

	nums[0] += "mahdi"
	max, _ := strconv.Atoi(nums2[0])
	var turn int
	for i := 1; i < len(nums2); i++ {
		turn, _ = strconv.Atoi(nums2[i])
		if turn > max {
			max = turn
		}
	}
	stringMax := strconv.Itoa(max)

	fmt.Println(stringMax)
}
