package main

import (
	"fmt"
	"strings"
)

/*
	problem link : https://quera.org/problemset/3430
*/

func main() {

	var word string
	fmt.Scan(&word)
	line := strings.Split(word, "")
	var avali, second, final string

	println(word)
	for i := 1; i < len(word); i++ {
		avali = strings.Repeat(line[i], i+1)
		second = word[i+1:]
		final = avali + second
		fmt.Println(final)
	}
}
