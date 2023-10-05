/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// package leetcodesolution

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	nqueue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		for i, node := range queue {
			if i < len(queue)-1 {
				node.Next = queue[i+1]
			}

			if node.Left != nil {
				nqueue = append(nqueue, node.Left)
			}
			if node.Right != nil {
				nqueue = append(nqueue, node.Right)
			}
		}

		queue = nqueue
		nqueue = make([]*Node, 0)
	}

	return root
}

// @lc code=end
