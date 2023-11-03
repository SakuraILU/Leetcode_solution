/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt)
}

func dfs(root *TreeNode, maxval int) (ngood int) {
	if root == nil {
		return 0
	}

	if root.Val >= maxval {
		ngood++
	}

	maxval = max(root.Val, maxval)
	ngood += dfs(root.Left, maxval) + dfs(root.Right, maxval)

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
