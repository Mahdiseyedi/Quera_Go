package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	problem link : https://quera.org/problemset/10230
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	result := strings.TrimSpace(input)
	nums := strings.Split(result, " ")

	var angel1, angel2, angel3 int
	var output = "Yes"
	angel1, _ = strconv.Atoi(nums[0])
	angel2, _ = strconv.Atoi(nums[1])
	angel3, _ = strconv.Atoi(nums[2])

	if angel1 == 0 || angel1 == 180 {
		output = "No"
	}
	if angel2 == 0 || angel2 == 180 {
		output = "No"
	}
	if angel3 == 0 || angel3 == 180 {
		output = "No"
	}
	if angel1+angel2+angel3 != 180 {
		output = "No"
	}
	fmt.Println(output)
}

func add(num int) (x, y int) {
	x = num * 2
	y = num - 3
	return
}
