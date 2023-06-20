package main

import (
	"fmt"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeDuplicates(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		if contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}

func loghat(name string) int {
	res := 0
	jname := strings.Split(name, "")
	jname = removeDuplicates(jname)
	res = len(jname)
	return res
}

func maxi(name []string) int {
	max := loghat(name[0])
	for i := 0; i < len(name); i++ {
		if loghat(name[i]) > max {
			max = loghat(name[i])
		}
	}
	return max
}

func main() {
	var n int
	var name string
	var names []string
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		names = append(names, name)
		name = ""
	}

	fmt.Println(maxi(names))
}
