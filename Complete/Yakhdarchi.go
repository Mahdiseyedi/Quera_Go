package main

import "fmt"

func main() {
	var temprature int
	fmt.Scanf("%d", &temprature)
	if temprature < 0 {
		fmt.Println("Ice")
	}
	if temprature <= 100 && temprature >= 0 {
		fmt.Println("Water")
	}
	if temprature > 100 {
		fmt.Println("Steam")
	}
}
