/*
 * @lc app=leetcode id=919 lang=golang
 *
 * [919] Complete Binary Tree Inserter
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// package leetcodesolution

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type CBTInserter struct {
	root      *TreeNode
	queue     []*TreeNode
	i         int
	new_queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	c := CBTInserter{
		root:      root,
		queue:     make([]*TreeNode, 0),
		i:         0,
		new_queue: make([]*TreeNode, 0),
	}

	c.queue = append(c.queue, root)
	for {
		c.new_queue = make([]*TreeNode, 0)
		for i, node := range c.queue {
			if node.Left != nil {
				c.new_queue = append(c.new_queue, node.Left)
			}
			if node.Right != nil {
				c.new_queue = append(c.new_queue, node.Right)
			}

			if node.Right == nil {
				c.i = i
				return c
			}
		}

		c.queue = c.new_queue
	}
}

func (this *CBTInserter) Insert(val int) int {
	node := this.queue[this.i]
	if node.Left == nil {
		node.Left = &TreeNode{Val: val}
		this.new_queue = append(this.new_queue, node.Left)
	} else if node.Right == nil {
		node.Right = &TreeNode{Val: val}
		this.new_queue = append(this.new_queue, node.Right)
		this.i++
	}

	if this.i >= len(this.queue) {
		this.queue = this.new_queue
		this.new_queue = make([]*TreeNode, 0)
		this.i = 0
	}
	return node.Val // return the parent of this inserted node
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
// @lc code=end
