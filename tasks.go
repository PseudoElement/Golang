package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isHappy(n int) bool {
	slice := make([]int, 0)
	return getPowsSum(&slice, n)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getPowsSum(slice *[]int, num int) bool {
	numStr := strconv.Itoa(num)
	chars := strings.Split(numStr, "")

	var sum int
	for _, char := range chars {
		digit, _ := strconv.Atoi(char)
		sum += digit * digit
	}

	if contains(*slice, sum) {
		return false
	}

	*slice = append(*slice, sum)

	if sum == 1 {
		return true
	}

	return getPowsSum(slice, sum)
}

func main() {
	x := isHappy(3)

	fmt.Println(x)
}
