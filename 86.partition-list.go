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
	slow := dummy

	for slow.Next != nil && slow.Next.Val < x {
		slow = slow.Next
	}
	fast := slow

	for fast.Next != nil {
		if fast.Next.Val < x {
			target := fast.Next
			fast.Next = target.Next

			target.Next = slow.Next
			slow.Next = target

			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}

	return dummy.Next
}

// @lc code=end
