/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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

var targetsum int = 0
var cursum int = 0
var has bool = false

func hasPathSum(root *TreeNode, targetSum int) bool {
	targetsum = targetSum
	cursum = 0
	has = false
	traverse(root)
	if has {
		return true
	} else {
		return false
	}
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	if has {
		return
	}

	cursum += node.Val

	if node.Left == nil && node.Right == nil {
		if cursum == targetsum {
			has = true
			return
		}
	}

	traverse(node.Left)
	traverse(node.Right)

	cursum -= node.Val
}

// @lc code=end
