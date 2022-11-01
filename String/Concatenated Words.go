package String

import (
	"sort"
)

//给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。
//
//连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。
//
//
//
//示例 1：
//
//输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
//输出：["catsdogcats","dogcatsdog","ratcatdogcat"]
//解释："catsdogcats" 由 "cats", "dog" 和 "cats" 组成;
//"dogcatsdog" 由 "dog", "cats" 和 "dog" 组成;
//"ratcatdogcat" 由 "rat", "cat", "dog" 和 "cat" 组成。
//示例 2：
//
//输入：words = ["cat","dog","catdog"]
//输出：["catdog"]
//
//
//提示：
//
//1 <= words.length <= 104
//0 <= words[i].length <= 30
//words[i] 仅由小写字母组成
//0 <= sum(words[i].length) <= 105

type trie struct {
	children [26]*trie //这是一个trie指针数组，26是因为最多26个字母
	isEnd    bool
}

func (root *trie) insert(word string) {
	node := root
	for _, ch := range word {
		//相减后再赋值ch=ch-'a',如果一开始就是a那么就是97，那么减去之后得到的ch就是children这个指针数组的位置是在26字母的第几个位置
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (root *trie) dfs(vis []bool, word string) bool {
	//空的单词，或者已经被剪没了
	if word == "" {
		return true
	}
	//对于每个单词，创建与单词相同长度的数组记录该单词的每一个下标是否被访问过，然后进行记忆化搜索。搜索过程中，如果一个下标已经被访问过，
	//则从该下标到末尾的部分一定不是由给定数组中的一个或多个非空单词组成（否则上次访问时已经可以知道当前单词是连接词），
	//只有尚未访问过的下标才需要进行搜索。

	//这个就说明下标被访问过了，如何理解呢，比如c cat catcat，那么catcat这个单词的最后一位没有被访问过，可以下一步
	//但是问题是，会出现这样的情况吗？
	if vis[len(word)-1] {
		return false
	}

	//标记这个下标被访问过
	vis[len(word)-1] = true

	node := root
	for i, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
		if node.isEnd && root.dfs(vis, word[i+1:]) {
			return true
		}
	}
	return false
}

func FindAllConcatenatedWordsInADict(words []string) (ans []string) {
	//这是sort包中Slice函数的用法，左边是排序的切片接口，右边是排序的规则，源码里面这是一个很精妙的设计
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })

	//root作为一个指针
	root := &trie{}
	for _, word := range words {

		//排除空单词
		if word == "" {
			continue
		}

		//这是单词每个字母的标记，
		vis := make([]bool, len(word))
		//如果是true，说明这个单词的的中途某个位置是某个较短单词的结尾就说明他是一个长的连接词
		if root.dfs(vis, word) {
			ans = append(ans, word)
		} else {
			//如果不是，哪怕前面有一部分是某个单词的结尾，那也要重新插入
			root.insert(word)
		}
	}
	return
}

//方法一：字典树 + 记忆化搜索
//判断一个单词是不是连接词，需要判断这个单词是否完全由至少两个给定数组中的更短的非空单词（可以重复）组成。判断更短的单词是否在给定数组中，可以使用字典树实现。
//
//为了方便处理，首先将数组 \textit{words}words 按照字符串的长度递增的顺序排序，排序后可以确保当遍历到任意单词时，比该单词短的单词一定都已经遍历过，因此可以根据已经遍历过的全部单词判断当前单词是不是连接词。
//
//在将数组 \textit{words}words 排序之后，遍历数组，跳过空字符串，对于每个非空单词，判断该单词是不是连接词，如果是连接词则将该单词加入结果数组，如果不是连接词则将该单词加入字典树。
//
//判断一个单词是不是连接词的做法是在字典树中搜索。从该单词的第一个字符（即下标 00 处的字符）开始，在字典树中依次搜索每个字符对应的结点，可能有以下几种情况：
//
//如果一个字符对应的结点是单词的结尾，则找到了一个更短的单词，从该字符的后一个字符开始搜索下一个更短的单词；
//
//如果一个字符对应的结点在字典树中不存在，则当前的搜索结果失败，回到上一个单词的结尾继续搜索。
//
//如果找到一个更短的单词且这个更短的单词的最后一个字符是当前单词的最后一个字符，则当前单词是连接词。由于数组 \textit{words}words 中没有重复的单词，因此在判断一个单词是不是连接词时，该单词一定没有加入字典树，由此可以确保判断连接词的条件成立。
//
//由于一个连接词由多个更短的非空单词组成，如果存在一个较长的连接词的组成部分之一是一个较短的连接词，则一定可以将这个较短的连接词换成多个更短的非空单词，因此不需要将连接词加入字典树。
//
//为了降低时间复杂度，需要使用记忆化搜索。对于每个单词，创建与单词相同长度的数组记录该单词的每一个下标是否被访问过，然后进行记忆化搜索。搜索过程中，如果一个下标已经被访问过，则从该下标到末尾的部分一定不是由给定数组中的一个或多个非空单词组成（否则上次访问时已经可以知道当前单词是连接词），只有尚未访问过的下标才需要进行搜索。
//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/concatenated-words/solution/lian-jie-ci-by-leetcode-solution-mj4d/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
