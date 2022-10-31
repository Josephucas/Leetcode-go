package leetcode

// LongestCommonPrefix 题目大意 #
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 “"。
//
//解题思路 #
//对 strs 按照字符串长度进行升序排序，求出 strs 中长度最小字符串的长度 minLen
//逐个比较长度最小字符串与其它字符串中的字符，如果不相等就返回 commonPrefix,否则就把该字符加入 commonPrefix
func LongestCommonPrefix(strs []string) string {
	prefix := strs[0]
	//比较第一个和后面的几个

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}
