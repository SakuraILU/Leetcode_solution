/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
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

var new_head *ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	traverse(head)

	head.Next = nil

	return new_head
}

func traverse(node *ListNode) {
	if node.Next == nil {
		new_head = node
		return
	}

	traverse(node.Next)

	node.Next.Next = node
}

// @lc code=end
