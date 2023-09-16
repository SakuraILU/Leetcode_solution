/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + dfs(root, targetSum)
}

// 以node为顶点的树中，等于targetsum的root-to-any的路径数量
func dfs(node *TreeNode, targetsum int) int {
	if node == nil {
		return 0
	}

	curnum := 0
	if node.Val == targetsum {
		curnum += 1
	}
	curnum += dfs(node.Left, targetsum-node.Val) + dfs(node.Right, targetsum-node.Val)

	return curnum
}

// @lc code=end
