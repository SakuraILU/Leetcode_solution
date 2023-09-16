/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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

var sum = 0
var cursum = 0

func sumNumbers(root *TreeNode) int {
	sum = 0
	cursum = 0

	traverse(root)

	return sum
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	cursum = cursum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		sum += cursum
	}

	traverse(node.Left)
	traverse(node.Right)
	cursum /= 10
}

// @lc code=end
