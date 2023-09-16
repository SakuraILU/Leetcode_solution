/*
 * @lc app=leetcode id=1367 lang=golang
 *
 * [1367] Linked List in Binary Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	if dfs(head, root) {
		return true
	} else {
		return isSubPath(head, root.Left) || isSubPath(head, root.Right)
	}
}

func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	} else if root == nil || head.Val != root.Val {
		// head != nil
		return false
	}

	return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
}

// @lc code=end
