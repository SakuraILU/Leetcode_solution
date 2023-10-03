/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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
func averageOfLevels(root *TreeNode) (res []float64) {
	res = make([]float64, 0)

	queue := make([]*TreeNode, 0)
	nextqueue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextqueue = make([]*TreeNode, 0)
		cnt := 0
		for _, node := range queue {
			cnt += node.Val
			if node.Left != nil {
				nextqueue = append(nextqueue, node.Left)
			}
			if node.Right != nil {
				nextqueue = append(nextqueue, node.Right)
			}
		}
		res = append(res, float64(cnt)/float64(len(queue)))
		queue = nextqueue
		nextqueue = make([]*TreeNode, 0)
	}

	return
}

// @lc code=end

