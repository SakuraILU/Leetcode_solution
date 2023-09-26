/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	lnode := lowestCommonAncestor(root.Left, p, q)
	rnode := lowestCommonAncestor(root.Right, p, q)

	if lnode != nil && rnode != nil {
		return root
	} else if lnode != nil {
		return lnode
	} else {
		return rnode
	}
}

// @lc code=end
