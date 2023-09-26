/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var successor1 *ListNode
var successor2 *ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseN(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	traverse(head, n)

	head.Next = successor2

	return successor1
}

func traverse(node *ListNode, n int) {
	if n == 1 {
		successor1 = node
		successor2 = node.Next
		return
	}

	traverse(node.Next, n-1)

	node.Next.Next = node

	return
}

// @lc code=end
