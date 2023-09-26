/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
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

const (
	NOTEND = iota
	END
)

func isCompleteTree(root *TreeNode) bool {
	state := NOTEND
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		new_queue := make([]*TreeNode, 0)
		for _, node := range queue {
			switch state {
			case NOTEND:
				if node.Right == nil {
					state = END
				} else if node.Left == nil {
					return false
				}
			case END:
				if node.Left != nil || node.Right != nil {
					return false
				}
			}

			if node.Left != nil {
				new_queue = append(new_queue, node.Left)
			}
			if node.Right != nil {
				new_queue = append(new_queue, node.Right)
			}
		}

		queue = new_queue
	}

	return true
}

// @lc code=end
