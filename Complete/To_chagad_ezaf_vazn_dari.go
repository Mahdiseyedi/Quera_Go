package main

import (
	"fmt"
)

func main() {

	var weight, height float32
	var name string
	fmt.Scan(&weight, &height)

	result := (weight) / (height * height)

	if result < 18.5 {
		name = "Underweight"
	} else if result < 25 && result >= 18.5 {
		name = "Normal"
	} else if result < 30 && result >= 25 {
		name = "Overweight"
	} else {
		name = "Obese"
	}
	fmt.Printf("%.2f", result)
	fmt.Print("\n", name)
}
