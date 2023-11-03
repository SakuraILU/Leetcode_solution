/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (balanced bool, height int) {
	if root == nil {
		return true, 0
	}

	lbalanced, lheight := dfs(root.Left)
	rbalanced, rheight := dfs(root.Right)
	if !lbalanced || !rbalanced {
		return false, -1
	}
	if abs(lheight-rheight) > 1 {
		return false, -1
	}

	balanced = true
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// @lc code=end
