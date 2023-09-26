/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	root_val := postorder[len(postorder)-1]
	root := &TreeNode{Val: root_val, Left: nil, Right: nil}

	root_idx_inorder := 0
	for idx, elem := range inorder {
		if elem == root_val {
			root_idx_inorder = idx
			break
		}
	}

	if root_idx_inorder > 0 {
		root.Left = buildTree(inorder[0:root_idx_inorder], postorder[0:root_idx_inorder])
	}
	if root_idx_inorder < len(inorder)-1 {
		root.Right = buildTree(inorder[root_idx_inorder+1:len(inorder)], postorder[root_idx_inorder:len(postorder)-1])
	}

	return root
}

// @lc code=end
