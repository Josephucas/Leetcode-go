package main

import (
	String "Leetcode-go/String"
	"fmt"
)

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//var templateArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//targets := 5
	//result := leetcode.TwoSum(templateArray, targets)
	//fmt.Println(leetcode.IntToRoman(49))

	//fmt.Println(Array.LongestCommonPrefix(strs))

	//string
	//strings := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	//trie := String.Constructor208()
	//trie.Insert("apple")
	//fmt.Println(trie.Search("apple"))
	//fmt.Println(trie.Search("app"))
	//fmt.Println(trie.StartsWith("app")) // 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false
	//trie.Insert("app")
	//fmt.Println(trie.Search("app"))

	words := []string{"cat", "dog", "catdog"}
	word := String.FindAllConcatenatedWordsInADict(words)
	fmt.Println(word)

}
