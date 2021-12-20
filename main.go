package main

import (
	"Leetcode-go/Array"
	"fmt"
)

func main() {
	var templateArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	targets := 5
	result := leetcode.TwoSum(templateArray, targets)
	fmt.Println(result)

}
