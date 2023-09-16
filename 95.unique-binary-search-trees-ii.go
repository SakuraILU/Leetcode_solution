/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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

func generateTrees(n int) []*TreeNode {
	return backtrack(1, n)
}

func backtrack(l, h int) (trees []*TreeNode) {
	trees = make([]*TreeNode, 0)

	if l > h {
		trees = append(trees, nil)
		return
	}

	for i := l; i <= h; i++ {
		ltrees := backtrack(l, i-1)
		rtrees := backtrack(i+1, h)

		for _, ltree := range ltrees {
			for _, rtree := range rtrees {
				tree := &TreeNode{
					Val:   i,
					Left:  ltree,
					Right: rtree,
				}
				trees = append(trees, tree)
			}
		}
	}

	return
}

// @lc code=end
