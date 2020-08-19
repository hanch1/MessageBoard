package util

import "unicode/utf8"

/*
  字符串过滤 （前缀树实现）
*/
type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func NewTrie() Trie {
	var r Trie
	r.Root = NewTrieNode()
	return r
}

func NewTrieNode() *TrieNode {
	n := new(TrieNode)
	n.Children = make(map[rune]*TrieNode)
	return n
}

func (this *Trie) Insert(txt string) {
	if len(txt) < 1 {
		return
	}
	node := this.Root
	key := []rune(txt)
	for i := 0; i < len(key); i++ {
		if _, exists := node.Children[key[i]]; !exists {
			node.Children[key[i]] = NewTrieNode()
		}
		node = node.Children[key[i]]
	}

	node.End = true
}

func (trie *Trie) Replace(content string) string {
	if len(content) < 1 {
		return content
	}
	node := trie.Root
	key := []rune(content)
	var tmp []rune = nil
	len := len(key)
	for i := 0; i < len; i++ {
		if _, exists := node.Children[key[i]]; exists {
			node = node.Children[key[i]]
			if node.End == true{
				tmp = key
				c, _ := utf8.DecodeRuneInString("*")
				tmp[i] = c
			}
			for j := i + 1; j < len; j++ {
				if _, exists := node.Children[key[j]]; exists {
					node = node.Children[key[j]]
					if node.End == true {
						if tmp == nil {
							tmp = key
						}
						// 修改屏蔽词
						for t := i; t <= j; t++ {
							c, _ := utf8.DecodeRuneInString("*")
							tmp[t] = c
						}
						i = j
						node = trie.Root
						break
					}
				}
			}
			node = trie.Root
		}
	}
	if tmp == nil {
		return content
	} else {
		return string(tmp)
	}
}
