/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// package leetcodesolution

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func swapPairs(head *ListNode) *ListNode {
	return dfs(head)
}

func dfs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head.Next
	nn := n.Next

	n.Next = head
	nn = dfs(nn)

	head.Next = nn

	return n
}

// @lc code=end
