/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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

func diameterOfBinaryTree(root *TreeNode) int {
	mdiam, _ := dfs(root)

	return mdiam - 1
}

func dfs(node *TreeNode) (mdiameter int, height int) {
	if node == nil {
		return 0, 0
	}

	lmdiam, lheight := dfs(node.Left)
	rmdiam, rheight := dfs(node.Right)

	mdiameter = max(max(lmdiam, rmdiam), lheight+rheight+1)
	height = max(lheight, rheight) + 1

	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
