/*
 * @lc app=leetcode id=894 lang=golang
 *
 * [894] All Possible Full Binary Trees
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

func allPossibleFBT(n int) (trees []*TreeNode) {
	if n < 1 {
		return
	} else if n == 1 {
		trees = append(trees, &TreeNode{Val: 0})
		return
	}

	ln := 1
	// generate subtrees with n - 1 nodes (root takes one)
	for ln < n-1 {
		ltrees := allPossibleFBT(ln)
		rtrees := allPossibleFBT((n - 1) - ln)

		for _, ltree := range ltrees {
			for _, rtree := range rtrees {
				tree := &TreeNode{
					Val:   0,
					Left:  ltree,
					Right: rtree,
				}
				trees = append(trees, tree)
			}
		}
		ln += 2
	}

	return
}

// @lc code=end
