/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	root_val := preorder[0]
	root := &TreeNode{Val: root_val, Left: nil, Right: nil}
	if len(preorder) == 1 {
		return root
	}

	root_idx_inorder := 0
	for idx, elem := range inorder {
		if elem == root_val {
			root_idx_inorder = idx
			break
		}
	}

	if root_idx_inorder > 0 {
		root.Left = buildTree(preorder[1:root_idx_inorder+1], inorder[0:root_idx_inorder])
	}
	if root_idx_inorder < len(preorder)-1 {
		root.Right = buildTree(preorder[root_idx_inorder+1:len(preorder)], inorder[root_idx_inorder+1:len(inorder)])
	}

	return root
}

// @lc code=end
