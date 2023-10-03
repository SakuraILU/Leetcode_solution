/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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

const EMPTY = -1

var preval int
var mindist int

func getMinimumDifference(root *TreeNode) int {
	preval = EMPTY
	mindist = math.MaxInt

	traverse(root)

	return mindist
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Left)
	if preval != EMPTY {
		dist := root.Val - preval
		mindist = min(mindist, dist)
	}
	preval = root.Val
	traverse(root.Right)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
