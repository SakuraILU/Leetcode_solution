/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1)
	leaves2 := getLeaves(root2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

func getLeaves(root *TreeNode) (leaves []int) {
	if root == nil {
		return []int{}
	}

	leaves = append(leaves, getLeaves(root.Left)...)

	if root.Left == nil && root.Right == nil {
		leaves = append(leaves, root.Val)
	}

	leaves = append(leaves, getLeaves(root.Right)...)

	return
}

// @lc code=end
