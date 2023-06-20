package main

import (
	"bufio"
	"fmt"
	"os"
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

func checker(initSlice []string) int {
	result := 0
	days := []string{"shanbe", "1shanbe", "2shanbe", "3shanbe", "4shanbe", "5shanbe", "jome"}
	for i := 0; i < len(initSlice); i++ {
		if contains(days, initSlice[i]) {
			result++
		}
	}

	return 7 - result
}

func main() {
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	result1 := strings.TrimSpace(input1)
	nums1 := strings.Split(result1, " ")

	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	result2 := strings.TrimSpace(input2)
	nums2 := strings.Split(result2, " ")

	reader3 := bufio.NewReader(os.Stdin)
	input3, _ := reader3.ReadString('\n')
	result3 := strings.TrimSpace(input3)
	nums3 := strings.Split(result3, " ")

	reader4 := bufio.NewReader(os.Stdin)
	input4, _ := reader4.ReadString('\n')
	result4 := strings.TrimSpace(input4)
	nums4 := strings.Split(result4, " ")

	reader5 := bufio.NewReader(os.Stdin)
	input5, _ := reader5.ReadString('\n')
	result5 := strings.TrimSpace(input5)
	nums5 := strings.Split(result5, " ")

	reader6 := bufio.NewReader(os.Stdin)
	input6, _ := reader6.ReadString('\n')
	result6 := strings.TrimSpace(input6)
	nums6 := strings.Split(result6, " ")

	var weekDays []string
	weekDays = append(weekDays, nums1...)
	weekDays = append(weekDays, nums3...)
	weekDays = append(weekDays, nums5...)
	weekDays = append(weekDays, nums2...)
	weekDays = append(weekDays, nums4...)
	weekDays = append(weekDays, nums6...)
	fmt.Println(checker(removeDuplicates(weekDays)))
}
