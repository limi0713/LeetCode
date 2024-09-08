package main

type WordDictionary struct {
	wordMap map[string]bool
	node    *Node
}

type Node struct {
	c map[byte]*Node
}

func NewNode() *Node {
	n := &Node{}
	n.c = make(map[byte]*Node)
	return n
}

func Constructor211() WordDictionary {
	d := WordDictionary{}
	d.wordMap = make(map[string]bool)
	d.node = NewNode()
	return d
}

func (this *WordDictionary) AddWord(word string) {
	this.wordMap[word] = true
	n := this.node
	for i := range word {
		ch := word[i]
		_, ok := n.c[ch]
		if !ok {
			n.c[ch] = NewNode()
		}
		n = n.c[ch]
	}
}

func (this *WordDictionary) Search(word string) bool {
	if this.wordMap[word] {
		return true
	}

	var dfs func(node *Node, i int) bool
	dfs = func(node *Node, i int) bool {
		if node == nil {
			return i == len(word)
		}
		if i >= len(word) {
			return false
		}

		if word[i] == '.' {
			for _, n := range node.c {
				if dfs(n, i+1) {
					return true
				}
			}
		} else if n, ok := node.c[word[i]]; ok {
			return dfs(n, i+1)
		}

		return false
	}

	return dfs(this.node, 0)
}
