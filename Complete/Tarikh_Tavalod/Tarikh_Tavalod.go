package main

import (
	"fmt"
	"strings"
)

/*
	problem link : https://quera.org/problemset/615
*/

func main() {
	var input string
	fmt.Scan(&input)

	tree := strings.Split(input, "")
	a := tree[0]
	b := tree[1]
	c := tree[2]
	d := tree[3]

	sal := a + b
	mah := c + d

	fmt.Print("saal:", sal, "\n")
	fmt.Print("maah:", mah)
}
