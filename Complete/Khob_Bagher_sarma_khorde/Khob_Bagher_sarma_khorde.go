package main

import (
	"fmt"
	"strings"
)

/*
	problem link : https://quera.org/problemset/10231
*/

func checker(s, t string) bool {
	result := false
	if strings.Contains(s, t) {
		result = true
	}
	return result
}

func main() {

	//, line3, line4, line5
	//, &line3, &line4, &line5
	var line1, line2, line3, line4, line5 string
	s1 := "MOLANA"
	s2 := "HAFEZ"
	tr := false
	final := ""

	fmt.Scan(&line1, &line2, &line3, &line4, &line5)
	if checker(line1, s1) || checker(line1, s2) {
		tr = false
		final += "1 "
	}
	if checker(line2, s1) || checker(line2, s2) {
		tr = false
		final += "2 "
	}
	if checker(line3, s1) || checker(line3, s2) {
		tr = false
		final += "3 "
	}
	if checker(line4, s1) || checker(line4, s2) {
		tr = false
		final += "4 "
	}
	if checker(line5, s1) || checker(line5, s2) {
		tr = false
		final += "5 "
	}
	if tr {
		final = "NOT FOUND!"
	}
	fmt.Println(final)
}
