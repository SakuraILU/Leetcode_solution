/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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

var prenode *TreeNode
var mismatch_x *TreeNode
var mismatch_y *TreeNode

func recoverTree(root *TreeNode) {
	prenode = nil
	mismatch_x, mismatch_y = nil, nil

	traverse(root)

	swap(mismatch_x, mismatch_y)

	return
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Left)
	// check order
	// error
	if prenode != nil && prenode.Val > root.Val {
		if mismatch_x == nil {
			// 两个mismatch的元素可能挨在一起，例如[1,3,2,4]
			// 这种情况只会导致一处位置违反中序单增的性质
			// 因此在未探测到第二处时，默认第二个元素为mismatch node
			mismatch_x = prenode
			mismatch_y = root
		} else {
			// 之后探测到第二处了，重置mismatch_y
			mismatch_y = root
		}
	}
	prenode = root

	traverse(root.Right)
}

func swap(node1, node2 *TreeNode) {
	tmpVal := node1.Val
	node1.Val = node2.Val
	node2.Val = tmpVal
}

// @lc code=end
