/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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

import "math"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var currank int = 0
var _k int = 0
var res int = 0

func kthSmallest(root *TreeNode, k int) int {
	currank = 0
	_k = k
	res = math.MaxInt

	traverse(root)

	return res
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	if res != math.MaxInt {
		return
	}

	traverse(node.Left)
	currank++
	if currank == _k {
		res = node.Val
	}
	traverse(node.Right)
}

// @lc code=end
