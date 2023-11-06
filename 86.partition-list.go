/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy

	for first.Next != nil && first.Next.Val < x {
		first = first.Next
	}
	second := first

	for second.Next != nil {
		if second.Next.Val < x {
			// move second.Next to first.Next
			target := second.Next
			second.Next = target.Next

			target.Next = first.Next
			first.Next = target

			first = first.Next
		} else {
			second = second.Next
		}
	}

	return dummy.Next
}

// @lc code=end
