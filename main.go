package main

import (
	leetcode "Leetcode-go/Array"
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	//var templateArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//targets := 5
	//result := leetcode.TwoSum(templateArray, targets)
	//fmt.Println(leetcode.IntToRoman(49))

	fmt.Println(leetcode.LongestCommonPrefix(strs))

}
