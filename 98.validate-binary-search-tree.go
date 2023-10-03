/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, l, h int) (isBST bool) {
	if root == nil {
		return true
	}

	if !(root.Val > l && root.Val < h) {
		return false
	}

	return dfs(root.Left, l, root.Val) && dfs(root.Right, root.Val, h)
}

// @lc code=end
