/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start

type Node struct {
	isword   bool
	children map[byte]*Node
}

func newNode() *Node {
	return &Node{
		isword:   false,
		children: make(map[byte]*Node),
	}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: newNode(),
	}
}

func (this *Trie) insert(node *Node, word string) {
	c := word[0]
	child, ok := node.children[c]
	if !ok {
		child = newNode()
		node.children[c] = child
	}

	if len(word) > 1 {
		this.insert(child, word[1:len(word)])
	} else {
		child.isword = true
	}
}

func (this *Trie) Insert(word string) {
	this.insert(this.root, word)
}

func (this *Trie) search_node(node *Node, str string) *Node {
	c := str[0]
	child, ok := node.children[c]
	if !ok {
		return nil
	}

	if len(str) > 1 {
		return this.search_node(child, str[1:len(str)])
	} else {
		return child
	}
}

func (this *Trie) Search(word string) bool {
	node := this.search_node(this.root, word)

	return node != nil && node.isword
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.search_node(this.root, prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
