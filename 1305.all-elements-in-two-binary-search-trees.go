/*
 * @lc app=leetcode id=1305 lang=golang
 *
 * [1305] All Elements in Two Binary Search Trees
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

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	vals1 := make([]int, 0)
	traverse(root1, &vals1)
	vals2 := make([]int, 0)
	traverse(root2, &vals2)
	return merge(vals1, vals2)
}

func traverse(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, vals)
	*vals = append(*vals, root.Val)
	traverse(root.Right, vals)
}

func merge(val1, val2 []int) (vals []int) {
	vals = make([]int, 0)

	i, j := 0, 0
	for i < len(val1) && j < len(val2) {
		if val1[i] <= val2[j] {
			vals = append(vals, val1[i])
			i++
		} else {
			vals = append(vals, val2[j])
			j++
		}
	}

	if i < len(val1) {
		vals = append(vals, val1[i:]...)
	} else {
		vals = append(vals, val2[j:]...)
	}

	return vals
}

// @lc code=end
