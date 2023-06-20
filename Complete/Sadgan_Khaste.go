package main

import (
	"fmt"
	"strconv"
	"strings"
)

func capy(num1, num2 string) string {

	sign := " < "
	var finalResult string
	var a, b, c, e, f, g int
	pNum1 := strings.Split(num1, "")
	pNum2 := strings.Split(num2, "")

	a, _ = strconv.Atoi(pNum1[0])
	b, _ = strconv.Atoi(pNum1[1])
	c, _ = strconv.Atoi(pNum1[2])

	e, _ = strconv.Atoi(pNum2[0])
	f, _ = strconv.Atoi(pNum2[1])
	g, _ = strconv.Atoi(pNum2[2])

	if c > g {
		finalResult += num2
		finalResult += sign
		finalResult += num1
	} else if g > c {
		finalResult += num1
		finalResult += sign
		finalResult += num2
	} else if b > f {
		finalResult += num2
		finalResult += sign
		finalResult += num1
	} else if f > b {
		finalResult += num1
		finalResult += sign
		finalResult += num2
	} else if a > e {
		finalResult += num2
		finalResult += sign
		finalResult += num1
	} else if e > a {
		finalResult += num1
		finalResult += sign
		finalResult += num2
	} else {
		sign = " = "
		finalResult += num1
		finalResult += sign
		finalResult += num2
	}

	return finalResult
}

func main() {

	var num1, num2 string
	fmt.Scan(&num1, &num2)

	fmt.Println(capy(num1, num2))
}
