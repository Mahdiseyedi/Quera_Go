package main

import (
	"fmt"
)

/*
	problem link : https://quera.org/problemset/2636
*/

func capy(shah, vazir, rokh, fil, asb, sarbaz int) (int, int, int, int, int, int) {

	shahA := 1
	vazirA := 1
	rokhA := 2
	filA := 2
	asbA := 2
	sarbazA := 8

	shahc := shahA - shah
	vazirc := vazirA - vazir
	rokhc := rokhA - rokh
	filc := filA - fil
	asbc := asbA - asb
	sarbazc := sarbazA - sarbaz

	return shahc, vazirc, rokhc, filc, asbc, sarbazc
}

func main() {
	var shah, vazir, rokh, fil, asb, sarbaz int
	fmt.Scan(&shah, &vazir, &rokh, &fil, &asb, &sarbaz)
	fmt.Println(capy(shah, vazir, rokh, fil, asb, sarbaz))
}
