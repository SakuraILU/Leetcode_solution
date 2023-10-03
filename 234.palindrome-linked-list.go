/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var left *ListNode
var ispalind bool

func isPalindrome(head *ListNode) bool {
	ispalind = true

	left = head
	traverse(head)

	return ispalind
}

func traverse(right *ListNode) {
	if right == nil {
		return
	}

	traverse(right.Next)

	if !ispalind {
		return
	}

	if left.Val == right.Val {
		left = left.Next
	} else {
		ispalind = false
	}
}

// @lc code=end
