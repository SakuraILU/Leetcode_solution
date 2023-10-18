/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

func maxPathSum(root *TreeNode) int {
	msum, _ := dfs(root)

	return msum
}

func dfs(node *TreeNode) (msum int, rootmsum int) {
	if node == nil {
		return math.MinInt, 0
	}

	lmsum, lrootmsum := dfs(node.Left)
	rmsum, rrootmsum := dfs(node.Right)

	msum = node.Val
	rootmsum = node.Val
	if lrootmsum > 0 {
		msum += lrootmsum
		rootmsum = max(rootmsum, node.Val+lrootmsum)
	}
	if rrootmsum > 0 {
		msum += rrootmsum
		rootmsum = max(rootmsum, node.Val+rrootmsum)
	}

	msum = max(max(lmsum, rmsum), msum)

	return msum, rootmsum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
