/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	isdeplica := false
	val := 0
	if head.Next != nil && head.Val == head.Next.Val {
		isdeplica = true
		val = head.Val
	}

	if isdeplica {
		for head != nil && head.Val == val {
			head = head.Next
		}
		return deleteDuplicates(head)
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}

// @lc code=end
