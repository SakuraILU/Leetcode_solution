/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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

var height int = 0
var views []int

func rightSideView(root *TreeNode) []int {
	height = 0
	views = make([]int, 0)

	traverse(root)

	return views
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	height++

	if height > len(views) {
		views = append(views, root.Val)
	}

	traverse(root.Right)
	traverse(root.Left)
	height--
}

// @lc code=end
