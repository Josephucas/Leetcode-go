package String

//208Implement a trie with insert, search, and startsWith methods.
//
//Example:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // returns true
//trie.search("app");     // returns false
//trie.startsWith("app"); // returns true
//trie.insert("app");
//trie.search("app");     // returns true
//Note:
//
//You may assume that all inputs are consist of lowercase letters a-z.
//All inputs are guaranteed to be non-empty strings.
//题目大意 #
//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
//解题思路 #
//要求实现一个 Trie 的数据结构，具有 insert, search, startsWith 三种操作
//这一题就是经典的 Trie 实现。本题的实现可以作为 Trie 的模板。

//有人在children里面直接装链表，我感觉还是装一个map比较好

type Trie struct {
	isWord   bool           //是否是一个单词，或者说，是否是结尾
	children map[rune]*Trie //rune是Go语言中一种特殊的数据类型,它是int32的别名,几乎在所有方面等同于int32,用于区分字符值和整数值。
}

// Constructor208 /** Initialize your data structure here. */
func Constructor208() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

// Insert /** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	parent := t
	for _, ch := range word {
		//如果在children的map里面存在word当中的ch单词，则ok为true，反之
		if child, ok := parent.children[ch]; ok {
			//在这个新建的原有的child的指针结构上继续检索
			parent = child
		} else {
			//新建一个叶子指针，在这个叶子下继续检索
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild //这是下一级的叶节点map的对应的是这个单词的下一级叶子
			parent = newChild              //层层替换
		}
	}
	parent.isWord = true //到最后的末尾单词了，所以是一个单词，标记为true。
}

// Search /** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	parent := t
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	parent := t
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
